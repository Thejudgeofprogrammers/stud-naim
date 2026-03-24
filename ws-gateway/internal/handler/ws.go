package handler

import (
	"encoding/json"

	"net/http"
	"ws-gateway/internal/hub"
	chat "ws-gateway/internal/service/chat"
	jwt_service "ws-gateway/internal/service/jwt"

	"github.com/gorilla/websocket"
)

type WSHandler struct {
	hub         *hub.Hub
	chatService *chat.ChatService
	jwtService  jwt_service.JWTService
}

func NewWSHandler(h *hub.Hub, chatSrv *chat.ChatService, jwtSrv jwt_service.JWTService) *WSHandler {
	return &WSHandler{
		hub:         h,
		chatService: chatSrv,
		jwtService:  jwtSrv,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *WSHandler) HandleWS(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "token required", 400)
		return
	}

	claims, err := h.jwtService.Parse(token)
	if err != nil {
		http.Error(w, "unauthorized", 401)
		return
	}

	userID := claims.UserID

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &hub.Client{
		UserID: userID,
		Conn:   conn,
		Send:   make(chan []byte, 256),
	}

	h.hub.Register <- client

	go h.writePump(client)
	go h.readPump(client)
}

func (h *WSHandler) readPump(c *hub.Client) {
	defer func() {
		h.hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		var meta struct {
			Type string `json:"type"`
			To   string `json:"to"`
		}

		if err := json.Unmarshal(msg, &meta); err != nil {
			continue
		}

		switch meta.Type {
		
		case "message":
			h.chatService.HandleMessage(c.UserID, msg)

		case "history":
			history := h.chatService.GetHistory(c.UserID, meta.To)
			data, _ := json.Marshal(map[string]interface{}{
				"type": "history",
				"data": history,
			})
			c.Send <- data
		
		case "chat_list":
			list := h.chatService.GetChatList(c.UserID)
			data, _ := json.Marshal(map[string]interface{}{
				"type": "chat_list",
				"data": list,
			})
			c.Send <- data
		}
	}
}

func (h *WSHandler) writePump(c *hub.Client) {
	defer c.Conn.Close()

	for msg := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}
