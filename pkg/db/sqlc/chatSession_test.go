package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSessionByNumber(t *testing.T) {
	user := createUserTest(t)
	createChatArg := CreateChatSessionParams{
		UserID:     user.UserID,
		Payload:    "{}",
		ActionFlow: "review",
	}
	chatCreated, err := testQuerier.CreateChatSession(context.TODO(), createChatArg)
	require.NoError(t, err)
	require.NotEmpty(t, chatCreated)

	chatFound, err := testQuerier.GetChatSessionsByUserPhone(context.TODO(), user.Phone)
	require.NoError(t, err)
	require.NotEmpty(t, chatFound)

	require.Equal(t, chatFound.ChatSessionID, chatCreated.ChatSessionID)

	err = testQuerier.DeleteChatSession(context.TODO(), chatFound.ChatSessionID)
	require.NoError(t, err)
}
