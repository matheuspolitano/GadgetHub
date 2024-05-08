package chat

import (
	"strconv"
	"testing"

	"github.com/matheuspolitano/GadgetHub/utils"
	"github.com/stretchr/testify/require"
)

func TestGenerateResponse(t *testing.T) {
	payload := make(map[string]string)
	action := &Action{
		Name:        utils.RandString(10),
		Description: utils.RandString(30),
		Payload: &Payload{
			Key: "1212",
		},
		Regex:    "Hello",
		Response: "How r y?",
	}

	message := "Hello World!"

	newAction, err := action.CheckMessage(message, payload)
	require.NoError(t, err)
	require.NotEmpty(t, newAction)

	value, ok := payload[action.Payload.Key]
	require.True(t, ok)
	require.Equal(t, message, value)
}

func TestGenerateResponseParse(t *testing.T) {
	payload := make(map[string]string)
	action := &Action{
		Name:        utils.RandString(10),
		Description: utils.RandString(30),
		Payload: &Payload{
			Key: "age",
		},
		Regex:    `^\s*(\d+)\s*$`,
		Response: "How old r y?",
	}

	message := strconv.Itoa(int(utils.RandNumber(0, 100)))

	newAction, err := action.CheckMessage(message, payload)
	require.NoError(t, err)
	require.NotEmpty(t, newAction)

	value, ok := payload[newAction.Payload.Key]
	require.True(t, ok)

	require.Equal(t, message, value)
}

func TestGenerateIf(t *testing.T) {
	payload := make(map[string]string)
	yesAnswer := "Cool, i love Brasil"

	action := &Action{
		Name:        utils.RandString(10),
		Description: "Question if is from Brasil",
		Payload: &Payload{
			Key: "age",
		},
		Regex:    `\b(?i)(yes|no)\b`,
		Response: "How are from Brazil",
		IfElse: &IfElse{
			Regex: `\b(?i)(yes)\b`,
			If: &Action{
				Response: yesAnswer,
			}, Else: &Action{
				Response: "Where are y from?",
			},
		},
	}

	message := "YES"

	newAction, err := action.CheckMessage(message, payload)
	require.NoError(t, err)
	require.NotEmpty(t, newAction)

	require.Equal(t, yesAnswer, newAction.Response)
}

func TestGenerateElse(t *testing.T) {
	payload := make(map[string]string)
	elseAnswer := "Where are y from?"

	action := &Action{
		Name:        utils.RandString(10),
		Description: "Question if is from Brasil",
		Payload: &Payload{
			Key: "age",
		},
		Regex:    `\b(?i)(yes|no)\b`,
		Response: "How are from Brazil",
		IfElse: &IfElse{
			Regex: `\b(?i)(yes)\b`,
			If: &Action{
				Response: utils.RandString(20),
			}, Else: &Action{
				Response: elseAnswer,
			},
		},
	}

	message := "No"

	newAction, err := action.CheckMessage(message, payload)
	require.NoError(t, err)
	require.NotEmpty(t, newAction)

	require.Equal(t, elseAnswer, newAction.Response)
}
