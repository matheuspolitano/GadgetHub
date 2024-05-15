package chat

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
)

type HandlerMessage struct {
	ChatSession  db.ChatSession
	LastAction   *Action
	LastMessage  db.ChatMessage
	Message      string
	PayloadValue map[string]string
	Store        db.Store
}

func NewHandler(from, msg string, store db.Store, template Template) (*HandlerMessage, error) {
	chatSession, err := store.GetChatSessionsByUserPhone(context.Background(), from)
	if err != nil {
		return nil, err
	}
	flow, err := template.GetFlow(FlowName(chatSession.ActionFlow))
	if err != nil {
		return nil, err
	}
	lastMessage, err := store.GetChatMessage(context.Background(), chatSession.LastMessageID.Int32)
	if err != nil {
		return nil, err
	}
	action, err := flow.GetAction(lastMessage.Action)
	if err != nil {
		return nil, err
	}

	payloadValue, err := JSONToMap(chatSession.Payload)
	if err != nil {
		return nil, err
	}
	return &HandlerMessage{
		ChatSession:  chatSession,
		LastAction:   action,
		LastMessage:  lastMessage,
		Message:      msg,
		PayloadValue: payloadValue,
		Store:        store,
	}, nil
}

func (hm *HandlerMessage) NextMessage() (string, error) {
	var msgResponse = ""
	var chatSessionUpdateParm db.UpdateChatSessionParams
	action, err := hm.LastAction.CheckMessage(hm.Message, hm.PayloadValue)
	if err != nil {
		return "", err
	}
	payloadString, err := MapToJSON(hm.PayloadValue)
	if err != nil {
		return "", err
	}
	if action.NextAction == "$end" {
		chatSessionUpdateParm = db.UpdateChatSessionParams{
			Payload:       pgtype.Text{Valid: true, String: payloadString},
			ClosedAt:      pgtype.Timestamptz{Valid: true, Time: time.Now()},
			ChatSessionID: hm.ChatSession.ChatSessionID,
		}
		msgResponse = action.Response

	} else {
		newMessageParm := db.CreateChatMessageParams{
			ChatSessionID:   hm.ChatSession.ChatSessionID,
			MessageReceived: pgtype.Text{Valid: true, String: hm.Message},
			MessageSent:     pgtype.Text{Valid: true, String: action.Response},
			Action:          action.NextAction,
			ReceivedAt:      pgtype.Timestamptz{Valid: true, Time: time.Now()},
			SentAt:          pgtype.Timestamptz{Valid: true, Time: time.Now()},
			MessageBeforeID: pgtype.Int4{Valid: true, Int32: hm.LastMessage.ChatMessageID},
		}
		newMessage, err := hm.Store.CreateChatMessage(context.TODO(), newMessageParm)
		if err != nil {
			return "", err
		}

		chatSessionUpdateParm = db.UpdateChatSessionParams{
			LastMessageID: pgtype.Int4{Valid: true, Int32: newMessage.ChatMessageID},
			Payload:       pgtype.Text{Valid: true, String: payloadString},
			ChatSessionID: hm.ChatSession.ChatSessionID,
		}
		msgResponse = action.Response
	}
	_, err = hm.Store.UpdateChatSession(context.TODO(), chatSessionUpdateParm)
	if err != nil {
		return "", err
	}

	return msgResponse, nil
}
