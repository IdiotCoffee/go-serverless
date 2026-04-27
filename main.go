package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/IdiotCoffee/go-serverless/pkg/handlers"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

const tableName = "LambdaGoUser"

func handler(req events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {
	switch req.HTTPMethod {
	case http.MethodGet:
		return handlers.GetUser(req, tableName, dynaClient)
	case http.MethodPost:
		return handlers.CreateUser(req, tableName, dynaClient)
	case http.MethodPut:
		return handlers.UpdateUser(req, tableName, dynaClient)
	case http.MethodDelete:
		return handlers.DeleteUser(req, tableName, dynaClient)
	default:
		return handlers.UnhandledMethod()
	}

}
