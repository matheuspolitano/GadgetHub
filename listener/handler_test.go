//go:build exclude
// +build exclude

package listener

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendMessage(t *testing.T) {
	err := webhookHandlerTest.sendMessage("328096297047007", "34674232412", "oisdfsd")
	require.NoError(t, err)
}
