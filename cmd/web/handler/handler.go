package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"stori-card-challenge-account/domain/account"
	"stori-card-challenge-account/domain/user"
	accountInfra "stori-card-challenge-account/internal/infrastructure/account"
	usecases "stori-card-challenge-account/internal/usecases/account"
	"stori-card-challenge-account/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	aws_config_path = "/var/task/aws_config.json"
)

type RequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// HandleAPIGatewayProxyRequest is the Lambda handler function.
func HandleAPIGatewayProxyRequest(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Read AWS configuration from JSON file
	config, err := utils.ReadAWSConfig(aws_config_path)
	if err != nil {
		fmt.Println("Error reading AWS config:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "broken!",
		}, nil
	}
	var requestBody RequestBody
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Wrong type of request",
		}, nil
	}
	msg := validateRequestModel(requestBody)
	if msg != "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       msg,
		}, nil
	}
	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.AWSRegion),
		Credentials: credentials.NewStaticCredentials(os.Getenv("aws_access_key"), os.Getenv("aws_secret_key"), ""),
	})
	if err != nil {
		log.Print("Error creating session:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "broken!",
		}, nil
	}
	dynamoClient := utils.CreateDBConnection(sess)
	log.Print("dynamo client", dynamoClient)
	log.Print("dynamo table: ", config.DynamoTable)
	saveAccountRepository := accountInfra.NewAccountDBRepository(dynamoClient, config.DynamoTable)
	saveAccountUsecase := usecases.NewSaveAccountUsecase(saveAccountRepository)

	saveAccountModel := createUserAccount(requestBody)

	err = saveAccountUsecase.Execute(ctx, saveAccountModel)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "user has been created",
	}, nil
}

func validateRequestModel(rb RequestBody) string {
	if rb.FirstName == "" {
		return "missing first name"
	}

	if rb.LastName == "" {
		return "missing last name"
	}
	return ""
}

func createUserAccount(rb RequestBody) *account.Account {
	usr := user.NewUser(rb.FirstName, rb.LastName)

	return account.NewAccountForUser(usr)

}
