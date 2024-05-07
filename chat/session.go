package chat

import (
	"time"

	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
)

type ChatMessage struct {
	ChatMessageID   int
	ChatSession     *ChatSession
	MessageReceived string
	MessageSent     string
	ReceivedAt      time.Time
	SentAt          time.Time
	action          *Action
}

type ChatSession struct {
	ChatSessionID int
	FlowAction    *Flow
	UserID        *db.User
	Payload       map[string]string
	OpenedAt      time.Time
	ClosedAt      time.Time
	LastMessage   *ChatMessage
}

func (manager *ManagerChat) NewChatSession(flowName FlowName, user *db.User) (*ChatSession, error) {
	flowAction, err := manager.template.GetFlow(flowName)
	if err != nil {
		return nil, err
	}
	payload := make(map[string]string)
	chatSession := &ChatSession{
		ChatSessionID: 1,
		FlowAction:    flowAction,
		Payload:       payload,
		OpenedAt:      time.Now(),
		UserID:        user,
	}
	var message = ""
	err = manager.messager.Send(message, user.Phone)
	if err != nil {
		return nil, err
	}
	primaryAction, err := flowAction.GetPrimaryAction()
	if err != nil {
		return nil, err
	}
	ChatMessage := &ChatMessage{
		ChatMessageID: 1,
		ChatSession:   chatSession,
		MessageSent:   message,
		SentAt:        time.Now(),
		action:        primaryAction,
	}
	chatSession.LastMessage = ChatMessage
	return chatSession, nil
}

// func (manager *ManagerChat) ReceiveMessage(message string, chatSession *ChatSession) error {

// }
