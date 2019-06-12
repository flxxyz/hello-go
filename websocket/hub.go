package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

type Hub struct {
	clients    []*Client
	register   chan *Client
	unregister chan *Client
	upgrader   *websocket.Upgrader
}

func NewHub() *Hub {
	hub := new(Hub)
	hub.clients = make([]*Client, 0)
	hub.register = make(chan *Client, 0)
	hub.unregister = make(chan *Client, 0)
	hub.upgrader = &websocket.Upgrader{
		// Allow all origins
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	go hub.Run()

	return hub
}

func (hub *Hub) Run() {
	fmt.Println("Hub.Run() ... ")
	for {
		select {
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		}
	}
}

func (hub *Hub) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := hub.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "could not upgrade", http.StatusInternalServerError)
		return
	}

	playerID, _ := uuid.NewV4()
	client := &Client{
		id:     playerID.String(),
		conn:   conn,
		packet: make(chan []byte),
		hub:    hub,
	}
	hub.register <- client

	go client.Read()
	go client.Write()
}

func (hub *Hub) Send(message interface{}, currClient *Client) {
	data, _ := json.Marshal(message)

	fmt.Printf("Send() data = %#v\n", string(data))
	currClient.packet <- data
}

func (hub *Hub) Broadcast(message interface{}, ignoreClient *Client) {
	data, _ := json.Marshal(message)
	for _, c := range hub.clients {
		if c != ignoreClient {
			c.packet <- data
		}
	}
}

func (hub *Hub) onMessage(data []byte, client *Client) {
	var msg Msg
	if json.Unmarshal(data, &msg) != nil {
		return
	}
	msg.UserId = client.id
	fmt.Printf("onMessage() msg = %#v\n", msg)
	hub.Send(msg, client)
}

func (hub *Hub) onConnect(client *Client) {
	log.Println("client connected: ", client.conn.RemoteAddr())
	hub.clients = append(hub.clients, client)
}

func (hub *Hub) onDisconnect(client *Client) {
	log.Println("client disconnected: ", client.conn.RemoteAddr())
	client.Close()

	i := -1
	for j, c := range hub.clients {
		if c.id == client.id {
			i = j
			break
		}
	}

	copy(hub.clients[i:], hub.clients[i+1:])
	hub.clients[len(hub.clients)-1] = nil
	hub.clients = hub.clients[:len(hub.clients)-1]
}
