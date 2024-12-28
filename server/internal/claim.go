package internal

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("token expired")

type Claim struct {
	Id        uuid.UUID
	UserId    int64
	CreatedAt time.Time
	ExpireAt  time.Time
}

func NewClaim(userId int64, duration time.Duration) (*Claim, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Claim{
		Id:        id,
		UserId:    userId,
		CreatedAt: time.Now(),
		ExpireAt:  time.Now().Add(duration),
	}, nil
}

func (c *Claim) Valid() error {
	if c.ExpireAt.Before(time.Now()) {
		return ErrExpiredToken
	}

	return nil
}
