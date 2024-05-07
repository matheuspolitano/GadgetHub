package chat

import (
	"testing"

	"github.com/matheuspolitano/GadgetHub/utils"
	"github.com/stretchr/testify/require"
)

func TestGenerateResponse(t *testing.T) {
	payload := make(map[string]string)
	action := &Action{
		Name:        utils.RandString(10),
		Description: utils.RandString(30),
		SavePayload: true,
		PayloadKey:  "test",
		Regex:       "Hello",
		Response:    "How r y?",
	}

	message := "Hello World!"

	newAction, err := action.CheckMessage(message, payload)
	require.NoError(t, err)
	require.NotEmpty(t, newAction)

	value, ok := payload[action.PayloadKey]
	require.True(t, ok)
	require.Equal(t, message, value)
}

func TestGenerateResponseParse(t *testing.T) {
	payload := make(map[string]string)
	action := &Action{
		Name:        utils.RandString(10),
		Description: utils.RandString(30),
		SavePayload: true,
		ParseFunc:   "parseInt",
		PayloadKey:  "test",
		Regex:       "12",
		Response:    "How r y?",
	}

	message := "12"

	newAction, err := action.CheckMessage(message, payload)
	require.NoError(t, err)
	require.NotEmpty(t, newAction)

	value, ok := payload[newAction.PayloadKey]
	require.True(t, ok)

	require.Equal(t, message, value)
}
