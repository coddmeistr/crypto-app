package ws

import (
	"encoding/json"
	"errors"
	"strings"

	app "github.com/maxim12233/crypto-app-server/crypto"
	bitfinex "github.com/maxim12233/crypto-app-server/crypto/bitfinex_sdk"
)

type Channel string

const (
	channelCandles Channel = "candles"
)

func (c *Client) handleSubscribe(channel Channel, params string) error {
	splitParams := strings.Split(params, ":")
	switch channel {
	case channelCandles:

		if len(splitParams) < 2 {
			return errors.New("Not enough params passed")
		}

		m, err := bitfinex.NewBitfinex()
		if err != nil {
			return err
		}
		rcv, done, err := m.SetRealOHLCVConnection(splitParams[0], splitParams[1])
		if err != nil {
			return err
		}
		sub := NewSubscription(func() {
		}, func() ([]byte, error) {
			data, ok := <-rcv
			if !ok {
				return []byte{}, errors.New("Channel is closed")
			}

			b, err := json.Marshal(data)
			if err != nil {
				return []byte{}, err
			}
			return b, nil

		}, func() {
			done <- struct{}{}
		})
		c.AddSubscriber(sub)

	default:
		return app.ErrUnknownChannel
	}

	return nil
}
