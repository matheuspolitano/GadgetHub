package tokenManager

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

	token, payload, err := JWTManager.GenerateToken(payloadArgument)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)
}

func TestCheckJWTToken(t *testing.T) {
	payloadArgument := PayloadParameters{
		UserID:   int(utils.RandNumber(0, 1000)),
		Role:     utils.RandString(10),
		Duration: time.Minute * 4,
	}
	secretKey := utils.RandString(40)
	JWTManager, err := NewJWTManager(secretKey)
	require.NoError(t, err)
	require.NotEmpty(t, JWTManager)

	token, payload, err := JWTManager.GenerateToken(payloadArgument)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	tokenPayload, err := JWTManager.CheckToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, tokenPayload)

	require.Equal(t, payload.ID, tokenPayload.ID)
	require.Equal(t, payload.UserID, tokenPayload.UserID)
	require.Equal(t, payload.Role, tokenPayload.Role)
	require.True(t, payload.ExpiredAt.Equal(tokenPayload.ExpiredAt))
	require.True(t, payload.IssuedAt.Equal(tokenPayload.IssuedAt))
}
