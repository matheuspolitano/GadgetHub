package token

import (
	"errors"
	"fmt"
)

var (
	ErrExpiredToken       = errors.New("token has expired")
	ErrInvalidToken       = errors.New("token is invalid")
	ErrExpiredSizeInvalid = fmt.Errorf("error: The secret key must be at least %d characters long. Please provide a longer key", minLen)
)
