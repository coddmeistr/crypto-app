package bitfinex

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

const (
	wshost = "wss://api-pub.bitfinex.com/ws/2"
)

type WSAction struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
	Key     string `json:"key"`
}

type IBitfinex interface {
	SetRealOHLCVConnection(timebase string, symbol string) (<-chan ResponseOHLCVUpdate, chan<- struct{}, error)
}

type Bitfinex struct{}

func NewBitfinex() (IBitfinex, error) {
	return &Bitfinex{}, nil
}

func readInfoMessage(msg []byte) error {
	var info ResponseInfo
	if err := json.NewDecoder(bytes.NewReader(msg)).Decode(&info); err != nil {
		return err
	}
	return nil
}

func readStatusMessage(msg []byte) (*ResponseOHLCVUpdate, error) {
	var status ResponseInfo
	if err := json.NewDecoder(bytes.NewReader(msg)).Decode(&status); err != nil {
		return nil, err
	}
	if status.Event == "subscribed" {
		return &ResponseOHLCVUpdate{
			Status:    "subscribed",
			ChannelId: status.ChanId,
		}, nil
	} else {
		return nil, errors.New("Failed to subscribe")
	}
}

func readSnapshotMessage(msg []byte) (*ResponseOHLCVUpdate, error) {
	var parsedData []interface{}
	err := json.Unmarshal(msg, &parsedData)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()

	channelID := parsedData[0].(float64)
	ohlcvData := parsedData[1].([]interface{})

	var ohlcvArray []OHLCVItem
	for _, v := range ohlcvData {
		var ohlcvItem OHLCVItem
		item := v.([]interface{})
		ohlcvItem.Timestamp = int64(item[0].(float64))
		ohlcvItem.Open = item[1].(float64)
		ohlcvItem.Close = item[2].(float64)
		ohlcvItem.High = item[3].(float64)
		ohlcvItem.Low = item[4].(float64)
		ohlcvItem.Volume = item[5].(float64)
		ohlcvArray = append(ohlcvArray, ohlcvItem)
	}

	return &ResponseOHLCVUpdate{ChannelId: int64(channelID), Snapshot: &ohlcvArray, Status: "snapshot"}, nil
}

func readUpdateMessage(msg []byte) (*ResponseOHLCVUpdate, error) {
	var parsedData []interface{}
	err := json.Unmarshal(msg, &parsedData)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()

	channelID := parsedData[0].(float64)

	ohlcvData, ok := parsedData[1].([]interface{})
	if !ok {
		// Heartbeat check
		if parsedData[1].(string) == "hb" {
			return &ResponseOHLCVUpdate{ChannelId: int64(channelID), Update: nil, Status: "update"}, nil
		}
		return nil, errors.New("Failed to parse data")
	}

	var ohlcvItem OHLCVItem
	ohlcvItem.Timestamp = int64(ohlcvData[0].(float64))
	ohlcvItem.Open = ohlcvData[1].(float64)
	ohlcvItem.Close = ohlcvData[2].(float64)
	ohlcvItem.High = ohlcvData[3].(float64)
	ohlcvItem.Low = ohlcvData[4].(float64)
	ohlcvItem.Volume = ohlcvData[5].(float64)

	return &ResponseOHLCVUpdate{ChannelId: int64(channelID), Update: &ohlcvItem, Status: "update"}, nil
}

func (c *Bitfinex) SetRealOHLCVConnection(timebase string, symbol string) (<-chan ResponseOHLCVUpdate, chan<- struct{}, error) {
	socket, _, err := websocket.DefaultDialer.Dial(wshost, nil)
	if err != nil {
		return nil, nil, err
	}

	action := WSAction{
		Event:   "subscribe",
		Channel: "candles",
		Key:     fmt.Sprintf("trade:%s:t%sUSD", timebase, symbol),
	}

	s, err := json.Marshal(action)
	if err != nil {
		return nil, nil, err
	}
	err = socket.WriteMessage(websocket.TextMessage, []byte(string(s)))
	if err != nil {
		return nil, nil, err
	}

	// Main gourouting that handling messages
	// When done channels gets some value then everything terminating
	done := make(chan struct{})
	write := make(chan ResponseOHLCVUpdate)
	msgCount := 0
	go func() {
		defer func() {
			err := socket.Close()
			if err != nil {
				log.Printf("error: %v", err)
			}
			close(write)
			fmt.Println("Exiting bitfinex ohlcv main gouruting")
		}()

		/* Main loop, where everything is structured via select statement
		   If getting done value, then terminating function
		   If getting message via helping gouruting above, then
		   handling this message and writing result to the write channel */
		for {

			// If reading from done channel, means that there are signal from the outside to terminate
			select {
			case <-done:
				return
			default:

				_, message, err := socket.ReadMessage()
				if err != nil {
					return
				}

				// Updates handler
				if msgCount > 2 {

					resp, err := readUpdateMessage(message)
					if err != nil {
						fmt.Println(err)
						return
					}
					if resp != nil {
						write <- *resp
					} else {
						fmt.Println("Nil response, closing connection...")
						return
					}

				} else if msgCount == 0 {

					// Start messages handler
					err := readInfoMessage(message)
					if err != nil {
						fmt.Println(err)
						return
					}
					msgCount += 1

				} else if msgCount == 1 {

					// Status messages handler
					resp, err := readStatusMessage(message)
					if err != nil {
						fmt.Println(err)
						write <- ResponseOHLCVUpdate{
							Status: "failed",
						}
						return
					}
					write <- *resp
					msgCount += 1

				} else if msgCount == 2 {

					// Snapshot messages handler
					resp, err := readSnapshotMessage(message)
					if err != nil {
						fmt.Println(err)
						return
					}
					if resp != nil {
						write <- *resp
					} else {
						fmt.Println("Nil response, closing connection...")
						return
					}
					msgCount += 1

				}
			}
		}
	}()

	return write, done, nil
}
