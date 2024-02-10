package account

import (
	usr "stori-card-challenge-account/domain/user"

	"time"

	"github.com/google/uuid"
)

type Status int

const (
	Overdrawn       Status = 1
	PositiveBalance Status = 2
	Inactive        Status = 3
)

var statusStrings = map[Status]string{
	Overdrawn:       "Overdrawn",
	PositiveBalance: "Positive Balance",
	Inactive:        "Inactive",
}

type Account struct {
	Id           string    `json:"id"`
	DateCreated  time.Time `json:"date_created"`
	Status       string    `json:"status"`
	TotalBalance float64   `json:"total_balance"`
	User         usr.User  `json:"user"`
}

func NewAccountForUser(user usr.User, status string, totalBalance float64) *Account {
	return &Account{
		Id:           uuid.New().String(),
		DateCreated:  time.Now().UTC(),
		Status:       status,
		TotalBalance: totalBalance,
		User:         user,
	}
}
