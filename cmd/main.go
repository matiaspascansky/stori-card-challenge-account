package main

import (
	"fmt"
	"stori-card-challenge-account/domain/account"
	"stori-card-challenge-account/domain/user"
)

func main() {

	user := user.NewMockIDUser("matias", "pascansky")

	acc := account.NewAccountForUser(user.ID)

	fmt.Print("hello", user.FirstName, "you have created an account with id: ", acc.Id, " and the status is: ", acc.Status)
}
