package token

import (
	"github.com/golang-jwt/jwt/v5"
)

// define minimal secret key length
const minLen = 32

var signMethod = jwt.SigningMethodHS256

type JWTManager struct {
	secretKey string
}

func NewJWTManager(secretKey string) (*JWTManager, error) {
	if len(secretKey) < 32 {
		return nil, ErrExpiredSizeInvalid
	}
	return &JWTManager{secretKey}, nil
}

func (manager *JWTManager) GenerateToken(payloadParameters PayloadParameters) (string, error) {
	payload, err := NewPayload(payloadParameters)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(signMethod, payload)
	jwtTokenStr, err := jwtToken.SignedString([]byte(manager.secretKey))
	return jwtTokenStr, err
}
