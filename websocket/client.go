package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	id     string
	conn   *websocket.Conn
	packet chan []byte
	hub    *Hub
}

func (c Client) Close() {
	c.conn.Close()
	close(c.packet)
}

func (c *Client) Read() {
	defer func() {
		c.hub.unregister <- c
	}()

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.hub.onMessage(data, c)
	}
}

func (c *Client) Write() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case data, ok := <-c.packet:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.WriteMessage(websocket.BinaryMessage, data)
		}
	}
}
