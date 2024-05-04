package token

import (
	"testing"
	"time"

	"github.com/matheuspolitano/GadgetHub/utils"
	"github.com/stretchr/testify/require"
)

func TestGenerateJWTToken(t *testing.T) {
	payloadArgument := PayloadParameters{
		UserID:   int(utils.RandNumber(0, 1000)),
		Role:     utils.RandString(10),
		Duration: time.Minute * 4,
	}
	secretKey := utils.RandString(40)
	JWTManager, err := NewJWTManager(secretKey)
	require.NoError(t, err)
	require.NotEmpty(t, JWTManager)

	token, err := JWTManager.GenerateToken(payloadArgument)
	require.NoError(t, err)
	require.NotEmpty(t, token)

}
