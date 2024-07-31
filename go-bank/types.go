package main

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type AuthToken struct {
	Id        int
	ExpiresAt time.Time
	jwt.RegisteredClaims
}

type CreateAccountResponse struct {
	token string
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(100000000)),
		Balance:   0,
		CreatedAt: time.Now().UTC(),
	}
}
