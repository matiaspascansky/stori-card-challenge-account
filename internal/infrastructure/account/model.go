package account

import (
	"stori-card-challenge-account/domain/account"
	"time"
)

type AccountDTO struct {
	Id            string    `json:"id"`
	DateCreated   time.Time `json:"date_created"`
	Status        string    `json:"status"`
	TotalBalance  float64   `json:"total_balance"`
	UserID        int64     `json:"user_id"`
	UserFirstName string    `json:"user_first_name"`
	UserLastName  string    `json:"user_last_name"`
}

func FromAccountToDTO(a *account.Account) *AccountDTO {
	return &AccountDTO{
		Id:            a.Id,
		DateCreated:   a.DateCreated,
		Status:        a.Status,
		TotalBalance:  a.TotalBalance,
		UserID:        a.User.ID,
		UserFirstName: a.User.FirstName,
		UserLastName:  a.User.LastName,
	}
}
