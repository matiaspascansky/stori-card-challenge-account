package account

import (
	"stori-card-challenge-account/utils"
	"time"
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
	UserId       int64     `json:"user_id"`
}

func NewAccountForUser(userId int64) Account {
	idGenerator := utils.NewAccountIDGenerator()
	return Account{
		Id:          idGenerator.GenerateID(),
		DateCreated: time.Now().UTC(),
		Status:      statusStrings[Inactive],
	}
}
