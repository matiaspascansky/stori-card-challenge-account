package main

import (
	"stori-card-challenge-account/cmd/web/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	/*user := user.NewMockIDUser("matias", "pascansky")

	acc := account.NewAccountForUser(user.ID)*/
	lambda.Start(handler.HandleAPIGatewayProxyRequest)

	//fmt.Print("hello", user.FirstName, "you have created an account with id: ", acc.Id, " and the status is: ", acc.Status)
}
