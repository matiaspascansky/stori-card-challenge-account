package account

import (
	"stori-card-challenge-account/utils"
	"time"
)

type Account struct {
	Id           string    `json:"id"`
	DateCreated  time.Time `json:"date_created"`
	Status       string    `json:"status"`
	TotalBalance float64   `json:"total_balance"`
	UserId       int64     `json:"user_id"`
}

func NewAccountForUser(userId int64) Account {
	idGenerator := utils.NewAccountIDGenerator()
	return Account{}
}
