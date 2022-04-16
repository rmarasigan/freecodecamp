package user

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/rmarasigan/go-serverless/pkg/validators"
)

var (
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord     = "failed to fetch record"
	ErrorInvalidUserData         = "invalid user data"
	ErrorInvalidEmail            = "invalid email"
	ErrorCouldNotMarshalItem     = "could not mashal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item"
	ErrorUserAlreadyExists       = "user.User already exists"
	ErrorUserDoesNotExist        = "user.User does not exist"
)

// User contains all the basic information of a user.
type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// FetchUser returns a single record of user based on email parameter passed.
func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	// Create a query based on the email parameter
	input := &dynamodb.GetItemInput{
		// `Key` is a required field. The key is the primary key that uniquely identifies each item in the table.
		// An AttributeValue represents the data for an attribute.
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				// Contains a name-pair value. The name is the data type, and the value is the data.
				S: aws.String(email),
			},
		},
		// Specifying the table name
		TableName: aws.String(tableName),
	}

	// Gets the item and returns a set of attributes for the item with the given primary key.
	result, err := dynaClient.GetItem(input)
	if err != nil {
		// `errors.New` returns an error that formats as the given text.
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	// `new` initialize a struct and returns a pointer to an instance of User struct
	item := new(User)

	// Unmarshal it into actual User struct which front-end can understand as a JSON
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	return item, nil
}

// FetchUsers returns the whole list of users.
func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*[]User, error) {
	// Getting all the users
	input := &dynamodb.ScanInput{
		// Specifying the table name
		TableName: aws.String(tableName),
	}

	// Returns one or more items and item attributes by accessing every item in a table or a secondary index.
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	// `new` initialize a struct and returns a pointer to an instance of []User struct
	item := new([]User)

	// Unmarshal it  into actual user which front-end can understand as a JSON
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	return item, nil
}

// CreateUser checks if the email is valid and if the user already exist. If not, it will create or insert new data to DynamoDB.
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var user User

	// Unmarshal the JSON and stores data to the User struct. If it returns an error while unmarshaling,
	// return ErrorInvaliduserData
	if err := json.Unmarshal([]byte(req.Body), &user); err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}

	// Checks if email is valid. If not, return an error message saying it's an invalid email
	if !validators.IsEmailValid(user.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}

	// Check if the user already exist
	currentUser, _ := FetchUser(user.Email, tableName, dynaClient)

	// If the user exist, return an error message saying that the user already exist
	if currentUser != nil && len(currentUser.Email) != 0 {
		return nil, errors.New(ErrorUserAlreadyExists)
	}

	// Converting the record to dynamodb.AttributeValue type
	av, err := dynamodbattribute.MarshalMap((user))
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	// Creating the data that you want to send to dynamoDB
	input := &dynamodb.PutItemInput{
		// Map of attribute name-value pairs, one for each attribute
		Item: av,
		// Specifying table name
		TableName: aws.String(tableName),
	}

	// Creates a new item
	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}

	return &user, nil
}

// UpdateUser checks if the user exist, if it does it will update the specific User information.
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var user User

	// Unmarshal the JSON and stores data to the User struct. If it returns an error while unmarshaling,
	// return ErrorInvalidEmail
	if err := json.Unmarshal([]byte(req.Body), &user); err != nil {
		return nil, errors.New(ErrorInvalidEmail)
	}

	// Checks if the user does not exist
	currentUser, _ := FetchUser(user.Email, tableName, dynaClient)

	// If the user email does not exist, return an error message saying that the user does not exist
	if currentUser != nil && len(currentUser.Email) == 0 {
		return nil, errors.New(ErrorUserDoesNotExist)
	}

	// Converting the record to dynamodb.AttributeValue type
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	// Creating the data that you want to send to dynamoDB
	input := &dynamodb.PutItemInput{
		// Map of attribute name-value pairs, one for each attribute
		Item: av,
		// Specifying table name
		TableName: aws.String(tableName),
	}

	// Replaces an old item with a new item
	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}

	return &user, nil
}

// DeleteUser deletes the user based on the email parameter being passed.
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) error {
	// Gets the request email parameter
	email := req.QueryStringParameters["email"]

	// Create a delete query based on the email parameter
	input := &dynamodb.DeleteItemInput{
		// `Key` is a required field. The key is the primary key that uniquely identifies each item in the table.
		// An AttributeValue represents the data for an attribute.
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				// Contains a name-pair value. The name is the data type, and the value is the data.
				S: aws.String(email),
			},
		},
		// Specifying table name
		TableName: aws.String(tableName),
	}

	// Deletes a single item in a table by primary key
	_, err := dynaClient.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorCouldNotDeleteItem)
	}

	return nil
}
