//go:build !feature1
// +build !feature1

package chat

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
	"github.com/matheuspolitano/GadgetHub/utils"
	"github.com/stretchr/testify/require"
)

func generateUserParams(phone string) db.CreateUserParams {
	return db.CreateUserParams{
		FirstName:    utils.RandString(10),
		LastName:     utils.RandString(10),
		Email:        utils.RandEmail(),
		HashPassword: utils.RandString(20),
		Phone:        phone,
		UserRole:     "admin",
	}
}
func createUserTest(t *testing.T, phone string) db.User {
	userParms := generateUserParams(phone)
	userCreated, err := testQuerier.CreateUser(context.TODO(), userParms)
	require.NoError(t, err)
	require.NotEmpty(t, userCreated)
	return userCreated
}
func TestNewHandler(t *testing.T) {
	phone := utils.RandPhone()
	flowName := "product_review"
	user := createUserTest(t, phone)

	createChatParm := db.CreateChatSessionParams{
		UserID:     user.UserID,
		Payload:    "{}",
		ActionFlow: flowName,
	}
	chatSession, err := testQuerier.CreateChatSession(context.Background(), createChatParm)
	require.NoError(t, err)
	require.NotEmpty(t, chatSession)

	createChatMessageParm := db.CreateChatMessageParams{
		MessageSent:   pgtype.Text{String: "Hey, Did you like rate your order experience?", Valid: true},
		SentAt:        pgtype.Timestamptz{Valid: true, Time: time.Now()},
		Action:        "want_rate",
		ChatSessionID: chatSession.ChatSessionID,
	}

	chatMessage, err := testQuerier.CreateChatMessage(context.Background(), createChatMessageParm)
	require.NoError(t, err)
	require.NotEmpty(t, chatMessage)

	updateChatSesssionParms := db.UpdateChatSessionParams{
		ChatSessionID: chatSession.ChatSessionID,
		LastMessageID: pgtype.Int4{Valid: true, Int32: chatMessage.ChatMessageID},
	}
	updatedChatSession, err := testQuerier.UpdateChatSession(context.Background(), updateChatSesssionParms)
	require.NoError(t, err)
	require.NotEmpty(t, updatedChatSession)

	handlerMessage, err := NewHandler(phone, "yes", testQuerier, *templateTest)
	require.NoError(t, err)
	require.NotEmpty(t, handlerMessage)

	msg, err := handlerMessage.NextMessage()
	require.NoError(t, err)
	require.NotEmpty(t, msg)

	flowTest, err := templateTest.GetFlow(FlowName(flowName))
	require.NoError(t, err)
	require.NotEmpty(t, flowTest)

	actionTest, err := flowTest.GetAction("want_rate")
	require.NoError(t, err)
	require.NotEmpty(t, actionTest)

	require.Equal(t, msg, actionTest.IfElse.If.Response)
}
func TestNewHandlerNo(t *testing.T) {
	phone := utils.RandPhone()
	flowName := "product_review"
	user := createUserTest(t, phone)

	createChatParm := db.CreateChatSessionParams{
		UserID:     user.UserID,
		Payload:    "{}",
		ActionFlow: flowName,
	}
	chatSession, err := testQuerier.CreateChatSession(context.Background(), createChatParm)
	require.NoError(t, err)
	require.NotEmpty(t, chatSession)

	createChatMessageParm := db.CreateChatMessageParams{
		MessageSent:   pgtype.Text{String: "Hey, Did you like rate your order experience?", Valid: true},
		SentAt:        pgtype.Timestamptz{Valid: true, Time: time.Now()},
		Action:        "want_rate",
		ChatSessionID: chatSession.ChatSessionID,
	}

	chatMessage, err := testQuerier.CreateChatMessage(context.Background(), createChatMessageParm)
	require.NoError(t, err)
	require.NotEmpty(t, chatMessage)

	updateChatSesssionParms := db.UpdateChatSessionParams{
		ChatSessionID: chatSession.ChatSessionID,
		LastMessageID: pgtype.Int4{Valid: true, Int32: chatMessage.ChatMessageID},
	}
	updatedChatSession, err := testQuerier.UpdateChatSession(context.Background(), updateChatSesssionParms)
	require.NoError(t, err)
	require.NotEmpty(t, updatedChatSession)

	handlerMessage, err := NewHandler(phone, "no", testQuerier, *templateTest)
	require.NoError(t, err)
	require.NotEmpty(t, handlerMessage)

	msg, err := handlerMessage.NextMessage()
	require.NoError(t, err)
	require.NotEmpty(t, msg)

	flowTest, err := templateTest.GetFlow(FlowName(flowName))
	require.NoError(t, err)
	require.NotEmpty(t, flowTest)

	actionTest, err := flowTest.GetAction("want_rate")
	require.NoError(t, err)
	require.NotEmpty(t, actionTest)

	require.Equal(t, msg, actionTest.IfElse.Else.Response)

	err = testQuerier.DeleteUser(context.TODO(), user.UserID)
}
