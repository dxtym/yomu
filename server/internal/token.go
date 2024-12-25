package internal

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 16

var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrInvalidKeySize    = errors.New("invalid key size")
	ErrInvalidsignMethod = errors.New("invalid sign method")
)

type Token struct {
	secretKey string
}

func NewToken(secretKey string) (*Token, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, ErrInvalidKeySize
	}

	return &Token{secretKey: secretKey}, nil
}

func (tm *Token) CreateToken(userId uint, duration time.Duration) (string, error) {
	claim, err := NewClaim(userId, duration)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(tm.secretKey))
}

func (tm *Token) VerifyToken(token string) (*Claim, error) {
	t, err := jwt.ParseWithClaims(token, &Claim{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidsignMethod
		}
		return []byte(tm.secretKey), nil
	})

	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok && errors.Is(v.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claim, ok := t.Claims.(*Claim)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claim, nil
}
