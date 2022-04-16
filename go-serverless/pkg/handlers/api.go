package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// apiResponse returns the API Gateway response in JSON format
func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	// Setting response header Content-Type as application/json
	response := events.APIGatewayProxyResponse{Headers: map[string]string{
		"Content-Type": "application/json",
	}}

	// Setting the HTTP status code
	response.StatusCode = status

	// Encodes body interface as JSON
	stringBody, _ := json.Marshal(body)

	// Converts []byte as string
	response.Body = string(stringBody)

	return &response, nil
}
