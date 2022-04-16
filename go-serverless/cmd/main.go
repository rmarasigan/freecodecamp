package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/rmarasigan/go-serverless/pkg/handlers"
)

var (
	// DynamoDBAPI provides an interface to enable mocking the dynamodb.DynamoDB service client's API operation, paginators, and waiters.
	dynaClient dynamodbiface.DynamoDBAPI
)

func main() {
	// Gets environment variable of AWS_REGION. It is important
	// because it is where your Lambda function will be deployed.
	region := os.Getenv("AWS_REGION")

	// Creates a new session
	awsSession, err := session.NewSession(&aws.Config{
		// Setting your AWS region and Enpoints
		Region: aws.String(region),
	})

	if err != nil {
		return
	}

	// Creates a DynamoDB client from just a session
	dynaClient = dynamodb.New(awsSession)

	// Will execute your lambda handler / function
	lambda.Start(handler)
}

const tableName = "go-serverless"

// handler is the method in your function code that processes events. When your function is invoked, Lambda runs the handler method.
// API Gateway events consist of a request that was routed to a Lambda function by API Gateway.
func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	// Specifies the integration's HTTP method type.
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUser(req, tableName, dynaClient)
	case "POST":
		return handlers.CreateUser(req, tableName, dynaClient)
	case "PUT":
		return handlers.UpdateUser(req, tableName, dynaClient)
	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynaClient)
	default:
		return handlers.UnhandledMethod()
	}
}
