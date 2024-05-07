package chat

import "time"

type ChatMessage struct {
	ChatMessageID   int
	UserID          int
	MessageReceived string
	MessageSent     string
	ReceivedAt      time.Time
	SentAt          time.Time
	CurrentAction   string
}

type ChatSession struct {
	ChatSessionID int
	FlowAction    string
	Payload       map[string]string
	CreatedAt     time.Time
	LastMessageID int
}
