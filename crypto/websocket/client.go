package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	app "github.com/maxim12233/crypto-app-server/crypto"
)

type Event string

const (
	eventSubscribe Event = "subscribe"
	eventClose     Event = "close"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout: time.Second * 5,
	Subprotocols:     []string{"json"},
	Error:            handleHandshakeError,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleHandshakeError(w http.ResponseWriter, r *http.Request, status int, reason error) {
	response, err := json.Marshal(fmt.Sprintf("Error: %s", reason))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(response)
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	subscription ISubscribtion
}

func (c *Client) AddSubscriber(sub ISubscribtion) {
	c.RemoveSubscriber()
	c.subscription = sub
	sub.Start()
}

func (c *Client) RemoveSubscriber() {
	if c.subscription == nil {
		return
	}
	c.subscription.End()
	c.subscription = nil
}

func (c *Client) RemoveSubscriberRude() {
	if c.subscription == nil {
		return
	}
	c.subscription = nil
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var act Action
		if err := json.Unmarshal(message, &act); err != nil {
			log.Printf("error: %v", err)
			c.send <- UnknownRequestResponse
			continue
		}

		switch act.Event {
		case eventSubscribe:
			if err := c.handleSubscribe(act.Channel, act.Params); err != nil {
				log.Printf("error: %v", err)
				if errors.Is(err, app.ErrUnknownChannel) {
					c.send <- UnknownChannelResponse
					continue
				} else {
					c.send <- UnknownInternalErrorResponse
					continue
				}
			}
		case eventClose:
			return
		default:
			c.send <- UnknownEventResponse
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) subscriptionPump() {
	for {
		if c.subscription == nil {
			continue
		}
		data, ok := c.subscription.GetMessage()
		if !ok {
			continue
		}
		c.send <- data
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
	go client.subscriptionPump()
}
