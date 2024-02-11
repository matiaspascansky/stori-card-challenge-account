package account

import (
	"stori-card-challenge-account/domain/account"
	"stori-card-challenge-account/domain/user"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFromAccountToDTO(t *testing.T) {
	u := user.User{
		ID:        1,
		FirstName: "Matias",
		LastName:  "Pascansky",
	}
	a := &account.Account{
		Id:           "testID",
		DateCreated:  time.Now(),
		Status:       "testStatus",
		TotalBalance: 1000.5,
		User:         u,
	}

	dto := FromAccountToDTO(a)

	assert.Equal(t, a.Id, dto.Id)
	assert.Equal(t, a.DateCreated, dto.DateCreated)
	assert.Equal(t, a.Status, dto.Status)
	assert.Equal(t, a.TotalBalance, dto.TotalBalance)
	assert.Equal(t, a.User.FirstName, dto.UserFirstName)
	assert.Equal(t, a.User.LastName, dto.UserLastName)
	assert.Equal(t, a.User.ID, dto.UserID)

}
