package chat

import (
	"errors"
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
	LastMessageID   int
}

type ChatSession struct {
	ChatSessionID int
	FlowAction    *Flow
	User          *db.User
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
		User:          user,
	}

	err = manager.messager.Send(flowAction.StartMessage, user.Phone)
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
		MessageSent:   flowAction.StartMessage,
		SentAt:        time.Now(),
		action:        primaryAction,
	}
	chatSession.LastMessage = ChatMessage
	return chatSession, nil
}

func (manager *ManagerChat) ReceiveMessage(message string, chatSession *ChatSession) error {
	action, err := chatSession.LastMessage.action.CheckMessage(message, chatSession.Payload)
	if err != nil {
		return err
	}
	if action.Response != "" {

		// CRIAR FUNCAO DE UPDATE Message colando a reposta do usuario
		manager.messager.Send(action.Response, chatSession.User.Phone)
	}
	//verificar se exist next se fechar chanal
	// se existir um next pegat ou criar uma nova ChatMessage item
	return errors.New("s")
}
