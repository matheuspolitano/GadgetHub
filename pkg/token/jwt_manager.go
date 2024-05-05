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

func (manager *JWTManager) GenerateToken(payloadParameters PayloadParameters) (string, *Payload, error) {
	payload, err := NewPayload(payloadParameters)
	if err != nil {
		return "", nil, err
	}
	jwtToken := jwt.NewWithClaims(signMethod, payload)
	jwtTokenStr, err := jwtToken.SignedString([]byte(manager.secretKey))
	return jwtTokenStr, payload, err
}

func (manager *JWTManager) CheckToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(manager.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil

}
