package service

import (
	"encoding/json"
	"sort"
	"time"
	"ws-gateway/internal/hub"
)

type Message struct {
	ID      string `json:"id"`
	ChatID  string `json:"chat_id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
	Time    int64  `json:"time"`
}

type Chat struct {
	ID      string
	User1   string
	User2   string
	Messages []Message
}

type ChatService struct {
	hub *hub.Hub

	chats map[string]*Chat
}

func getChatID(u1, u2 string) string {
	if u1 < u2 {
		return u1 + ":" + u2
	}
	return u2 + ":" + u1
}

func NewChatService(h *hub.Hub) *ChatService {
	return &ChatService{
		hub:   h,
		chats: make(map[string]*Chat),
	}
}

func (s *ChatService) HandleMessage(from string, raw []byte) {
	var input struct {
		Type    string `json:"type"`
		To      string `json:"to"`
		Content string `json:"content"`
	}

	if err := json.Unmarshal(raw, &input); err != nil {
		return
	}

	chatID := getChatID(from, input.To)

	chat, ok := s.chats[chatID]
	if !ok {
		chat = &Chat{
			ID:    chatID,
			User1: from,
			User2: input.To,
		}
	}

	msg := Message{
		ID:      generateID(),
		ChatID:  chatID,
		From:    from,
		To:      input.To,
		Content: input.Content,
		Time:    time.Now().Unix(),
	}

	chat.Messages = append(chat.Messages, msg)

	data, _ := json.Marshal(msg)

	s.hub.SendTo(input.To, data)
	s.hub.SendTo(from, data)
}

type ChatPreview struct {
	UserID      string `json:"user_id"`
	LastMessage string `json:"last_message"`
	Time        int64  `json:"time"`
}

func (s *ChatService) GetChatList(userID string) []ChatPreview {
	var result []ChatPreview

	for _, chat := range s.chats {
		var other string

		if chat.User1 == userID {
			other = chat.User2
		} else if chat.User2 == userID {
			other = chat.User1
		} else {
			continue
		}

		if len(chat.Messages) == 0 {
			continue
		}

		last := chat.Messages[len(chat.Messages)-1]

		result = append(result, ChatPreview{
			UserID:      other,
			LastMessage: last.Content,
			Time:        last.Time,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Time > result[j].Time
	})

	return result
}

func (s *ChatService) GetHistory(user1, user2 string) []Message {
	chatID := getChatID(user1, user2)

	chat, ok := s.chats[chatID]
	if !ok {
		return []Message{}
	}

	return chat.Messages
}

func generateID() string {
	return time.Now().Format("20060102150405.000000")
}
