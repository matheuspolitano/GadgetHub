package tokenManager

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Payload load the data from token
type Payload struct {
	ID        uuid.UUID `json: "id"`
	UserID    int       `json: "user_id"`
	Role      string    `json: "role"`
	IssuedAt  time.Time `json: "issued_at"`
	ExpiredAt time.Time `json: "expired_at"`
}

type PayloadParameters struct {
	UserID   int
	Role     string
	Duration time.Duration
}

func NewPayload(payloadParameters PayloadParameters) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		UserID:    payloadParameters.UserID,
		Role:      payloadParameters.Role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(payloadParameters.Duration),
	}
	return payload, nil
}

func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	numericData := &jwt.NumericDate{
		Time: p.ExpiredAt,
	}
	return numericData, nil
}

func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	numericData := &jwt.NumericDate{
		Time: p.IssuedAt,
	}
	return numericData, nil
}

func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	numericData := &jwt.NumericDate{
		Time: p.IssuedAt,
	}
	return numericData, nil
}

func (p *Payload) GetIssuer() (string, error) {
	return p.ID.String(), nil
}

func (p *Payload) GetSubject() (string, error) {
	return strconv.Itoa(p.UserID), nil
}

func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return []string{p.Role}, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
