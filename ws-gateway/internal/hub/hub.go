package hub

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	UserID string
	Conn   *websocket.Conn
	Send   chan []byte
}

type Hub struct {
	clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client.UserID] = client
			log.Println("connected:", client.UserID)
		case client := <-h.Unregister:
			delete(h.clients, client.UserID)
			close(client.Send)
			log.Println("disconnect:", client.UserID)
		}
	}
}

func (h *Hub) SendTo(userID string, msg []byte) {
	if client, ok := h.clients[userID]; ok {
		select {
		case client.Send <- msg:
		default:
			close(client.Send)
			delete(h.clients, userID)
		}
	}
}
