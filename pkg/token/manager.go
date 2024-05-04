package token

import "time"

type TokenManager interface {
	GenerateToken(username string, role string, duration time.Duration) (string, error)
	CheckToken(token string) (*Payload, error)
}
