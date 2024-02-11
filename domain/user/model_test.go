package user

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	firstName := "Matias"
	lastName := "Pascansky"

	u := NewUser(firstName, lastName)

	idStr := strconv.FormatInt(u.ID, 10)

	assert.Equal(t, firstName, u.FirstName)
	assert.Equal(t, lastName, u.LastName)

	assert.Equal(t, 6, len(idStr))

}
