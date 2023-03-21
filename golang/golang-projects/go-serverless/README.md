# go-serverless

![CRUD Serverless](/assets/img/go-serverless.png "go-serverless")

## File Hierarchy
```
go-serverless
│   README.md
│   go.mod
│   go.sum
│
└───build
│   main.zip
|
└───cmd
│   main.go
│   
└───pkg
│   │
│   └───handlers
|   │   api.go
|   │   handlers.go
|   |
|   └───user
|   │   user.go
|   │
|   └───validators
|   │   iEmailValid.go
```

## Development
```bash
dev@dev:~/go/src/github.com/development$ mkdir go-serverless && cd go-serverless
dev@dev:~/go/src/github.com/development/go-serverless$ go mod init
go: creating new go.mod: module github.com/development/go-serverless
dev@dev:~/go/src/github.com/development/go-serverless$ go mod tidy
dev@dev:~/go/src/github.com/development/go-serverless$ cd cmd
dev@dev:~/go/src/github.com/development/go-serverless/cmd$ go build main.go
# Move your executable file to `build` folder to create a zip file that will be uploaded to your Lambda function.
dev@dev:~/go/src/github.com/development/go-serverless/cmd$ cd ..
dev@dev:~/go/src/github.com/development/go-serverless$ zip -jrm build/main.zip build/main
  adding: main (deflated 58%)
```

### Creating Lambda function
* Change your **Runtime** into **Go1.x**
![go-serverless-lambda](/assets/img/go-serverless-lambda.png "lambda-creation")

* Change the execution role and set the role name and policy template.
![go-serverless-lambda-execution-role](/assets/img/go-serverless-lambda-execution-role.png "lambda-execution-role")


* Edit runtime settings and change the Handler name to **main**.
![go-serverless-edit-runtime](/assets/img/go-serverless-edit-runtime.png "lambda-edit-runtime")

* Under the **Code** tab, upload your .zip file.
![go-serverless-upload-zip-file](/assets/img/go-serverless-upload-zip-file.png "lambda-upload-zip-file")

### DynamoDB
* Create the table
![go-serverless-dynamodb-create-table](/assets/img/go-serverless-dynamodb-create-table.png "dynamodb-create-table")


### Configuring API Gateway
* Create a REST API
![go-serverless-apigateway-create-rest-api](/assets/img/go-serverless-apigateway-create-rest.png "apigateway-create-rest-api")

* Create a **Method** and select **ANY**. Select the **Lambda Function** as your Integration Type and **use Lambda Proxy integration**.
![go-serverless-apigateway-create-method](/assets/img/go-serverless-apigateway-create-method.png "apigateway-create-method")

* Deploy API
![go-serverless-apigateway-deploy](/assets/img/go-serverless-apigateway-deploy.png "apigateway-deploy")

## Test API
```bash
# Insert new user
dev@dev:~/go/src/github.com/development/go-serverless$ curl --header "Content-Type: application/json" --request POST --data '{"email": "alex.holland@gmail.com", "firstName": "Alex", "lastName": "Holland"}' https://xxxxxxxxxxxxx.execute.api.xx.region.xx.amazon.aws.com/staging
{"email": "alex.holland@gmail.com", "firstName": "Alex", "lastName": "Holland"}
# Fetch all users
dev@dev:~/go/src/github.com/development/go-serverless$ curl -X GET https://xxxxxxxxxxxxx.execute.api.xx.region.xx.amazon.aws.com/staging
[{"email": "alex.holland@gmail.com", "firstName": "Alex", "lastName": "Holland"}]
# Fetch specific user
dev@dev:~/go/src/github.com/development/go-serverless$ curl -X GET https://xxxxxxxxxxxxx.execute.api.xx.region.xx.amazon.aws.com/staging\?email\=alex.holland@gmail.com
[{"email": "alex.holland@gmail.com", "firstName": "Alex", "lastName": "Holland"}]
# Update specific user
dev@dev:~/go/src/github.com/development/go-serverless$ curl --header "Content-Type: application/json" --request PUT --data '{"email": "alex.holland@gmail.com", "firstName": "lalala", "lastName": "blablabla"}' https://xxxxxxxxxxxxx.execute.api.xx.region.xx.amazon.aws.com/staging
[{"email": "alex.holland@gmail.com", "firstName": "lalala", "lastName": "blablabla"}]
# Delete specific user
dev@dev:~/go/src/github.com/development/go-serverless$ curl -X DELETE https://xxxxxxxxxxxxx.execute.api.xx.region.xx.amazon.aws.com/staging\?email\=alex.holland@gmail.com
null
dev@dev:~/go/src/github.com/development/go-serverless$ curl -X GET https://xxxxxxxxxxxxx.execute.api.xx.region.xx.amazon.aws.com/staging
[]
```



## Reference
* [DynamoDB, expressions and Go](https://antklim.medium.com/dynamodb-expressions-and-go-b8230c253e1f)
* [aws-lambda-go/events/](https://github.com/aws/aws-lambda-go/tree/main/events)