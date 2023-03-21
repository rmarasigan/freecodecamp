# lambda-basics

![AWS Lambda with Golang](/assets/img/aws-lambda-with-golang.png "aws-lambda-go")

Prerequisite:
- Background knowledge golang basics
- AWS Knowledge
- AWS CLI already configured

```bash
dev@dev:~/go/src/github.com/development/lambda-basics$ go mod init
go: creating new go.mod: module github.com/development/freecodecamp/lambda-basics
dev@dev:~/go/src/github.com/development/lambda-basics$ aws iam create-role --role-name lambda-ex --assume-role-policy-document '{"Version": "2012-10-17", "Statement": [{"Effect": "Allow", "Principal": {"Service": "lambda.amazonaws.com"}, "Action": "sts:AssumeRole"}]}'
{
    "Role": {
        "Path": "/",
        "RoleName": "lambda-ex",
        "RoleId": "BEYSCP3YERPA1BUG8FOZA",
        "Arn": "arn:aws:iam::419784325734:role/lambda-ex",
        "CreateDate": "2022-04-15T13:44:10+00:00",
        "AssumeRolePolicyDocument": {
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Effect": "Allow",
                    "Principal": {
                        "Service": "lambda.amazonaws.com"
                    },
                    "Action": "sts:AssumeRole"
                }
            ]
        }
    }
}
# You can create a file called, trust-policy.json, and do this command. As I already done the first command, it will say that it already exist.
dev@dev:~/go/src/github.com/development/lambda-basics$ aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
An error occurred (EntityAlreadyExists) when calling the CreateRole operation: Role with name lambda-ex already exists.
dev@dev:~/go/src/github.com/development/lambda-basics$ aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
dev@dev:~/go/src/github.com/development/lambda-basics$ go build main.go
dev@dev:~/go/src/github.com/development/lambda-basics$ ls
go.mod  go.sum  main  main.go  trust-policy.json
# You need your file to be in a zip file
dev@dev:~/go/src/github.com/development/lambda-basics$ zip function.zip main
  adding: main (deflated 46%)
dev@dev:~/go/src/github.com/development/lambda-basics$ ls
function.zip  go.mod  go.sum  main  main.go  trust-policy.json
# handler is basically main, inside that zip file is the main executable file
# please remember to change the arn (amazon resource name) to your arn
dev@dev:~/go/src/github.com/development/lambda-basics$ aws lambda create-function --function-name go-lambda-basics \
> --zip-file fileb://function.zip --handler main \
> --runtime go1.x --role arn:aws:iam::419784325734/lambda-ex
{
    "FunctionName": "go-lambda-basics",
    "FunctionArn": "arn:aws:lambda:us-east-1:419784325734:function:go-lambda-basics",
    "Runtime": "go1.x",
    "Role": "arn:aws:iam::419784325734:role/lambda-ex",
    "Handler": "main",
    "CodeSize": 1531507,
    "Description": "",
    "Timeout": 3,
    "MemorySize": 128,
    "LastModified": "2022-04-15T14:00:04.725+0000",
    "CodeSha256": "/uP5wERtsfHFRVfkxo4lHSPOIXa5sREHM1JPM6xeMPJ=",
    "Version": "$LATEST",
    "TracingConfig": {
        "Mode": "PassThrough"
    },
    "RevisionId": "121er6b3-u456-9923-aw2d-2c292epo1140",
    "State": "Pending",
    "StateReason": "The function is being created.",
    "StateReasonCode": "Creating",
    "PackageType": "Zip",
    "Architectures": [
        "x86_64"
    ]
}
# This will invoke your lambda function
dev@dev:~/go/src/github.com/development/lambda-basics$ aws lambda invoke --function-name go-lambda-basics --cli-binary-format raw-in-base64-out --payload '{"What is your name?": "Jim", "How old are you?": 33}' output.txt
{
    "StatusCode": 200,
    "ExecutedVersion": "$LATEST"
}
```