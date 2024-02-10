package main

import (
	"stori-card-challenge-account/cmd/web/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	//lambda.Start(handler.HandleAPIGatewayProxyRequest)
	lambda.Start(handler.HandleSNS)

}
