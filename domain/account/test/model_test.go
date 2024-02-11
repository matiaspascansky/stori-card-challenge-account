package account

import (
	"stori-card-challenge-account/domain/account"
	"stori-card-challenge-account/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccountForUser(t *testing.T) {
	u := user.User{
		ID:        1,
		FirstName: "Matias",
		LastName:  "Pascansky",
	}

	totalBalance := 1000.00
	status := "debt-free"

	acc := account.NewAccountForUser(u, status, float64(totalBalance))

	assert.Equal(t, u.ID, acc.User.ID)
	assert.Equal(t, u.FirstName, acc.User.FirstName)
	assert.Equal(t, u.LastName, acc.User.LastName)
	assert.Equal(t, status, acc.Status)
	assert.Equal(t, totalBalance, acc.TotalBalance)

}
