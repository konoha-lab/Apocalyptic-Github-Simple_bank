package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Issuer    string    `json:"issuer"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type PayloadFooter struct {
	Info string `json:"info"`
}

var (
	ErrInvalidToken         = errors.New("token is invalid")
	ErrExpiredToken         = errors.New("token has expired")
	APPLICATION_NAME string = "SIMPLE BANK"
)

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		Issuer:    APPLICATION_NAME,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func NewPayloadFooter() (*PayloadFooter, error) {
	footer := &PayloadFooter{
		Info: "Nothing to say",
	}

	return footer, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
