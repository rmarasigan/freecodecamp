# go-fiber-mongo-hrms

## Development
```bash
dev@dev:~$ cd go/src/github.com/development
dev@dev:~/go/src/github.com/development$ mkdir go-fiber-mongo-hrms && cd go-fiber-mongo-hrms
dev@dev:~/go/src/github.com/development/go-fiber-mongo-hrms$ go mod init
go: creating new go.mod: module github.com/development/go-fiber-mongo-hrms
dev@dev:~/go/src/github.com/development/go-fiber-mongo-hrms$ go get go.mongodb.org/mongo-driver/mongo
go: downloading go.mongodb.org/mongo-driver v1.9.0
go: downloading github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d
go: downloading golang.org/x/crypto v0.0.0-20220214200702-86341886e292
go: downloading golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
go: downloading github.com/xdg-go/scram v1.0.2
go: downloading github.com/xdg-go/stringprep v1.0.2
go: downloading github.com/go-stack/stack v1.8.0
go: downloading github.com/golang/snappy v0.0.1
go: downloading github.com/xdg-go/pbkdf2 v1.0.0
go get: added github.com/go-stack/stack v1.8.0
go get: added github.com/golang/snappy v0.0.1
go get: added github.com/pkg/errors v0.9.1
go get: added github.com/xdg-go/pbkdf2 v1.0.0
go get: added github.com/xdg-go/scram v1.0.2
go get: added github.com/xdg-go/stringprep v1.0.2
go get: added github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d
go get: added go.mongodb.org/mongo-driver v1.9.0
go get: added golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
dev@dev:~/go/src/github.com/development/go-fiber-mongo-hrms$ go mod tidy
```

## Build Application
```bash
dev@dev:~/go/src/github.com/development/go-fiber-mongo-hrms$ go build main.go
dev@dev:~/go/src/github.com/development/go-fiber-mongo-hrms$ go run main.go

 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.32.0                   │ 
 │               http://127.0.0.1:3000               │ 
 │       (bound on host 0.0.0.0 and port 3000)       │ 
 │                                                   │ 
 │ Handlers ............. 5  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............ 266106 │ 
 └───────────────────────────────────────────────────┘ 
 # Add new employee
 dev@dev:~/go/src/github.com/development/go-fiber-mongo-hrms$ curl -X POST -d '{"name": "rupert", "salary": 3424, "age": 24}' -H 'Content-Type: application/json' http://localhost:3000/employee
 {"id":"61d7f9d64291ec8f94b2ee43","name":"rupert","salary":3424,"age":24}
 dev@dev:~/go/src/github.com/development/go-fiber-mongo-hrms$ curl -X POST -d '{"name": "alex", "salary": 3424, "age": 39}' -H 'Content-Type: application/json' http://localhost:3000/employee
 {"id":"63w7f34d654291esdf94b2oa43","name":"rupert","salary":3424,"age":29}
 # Fetch all employees
 dev@dev:~/go/src/github.com/development/go-fiber-mongo-hrms$ curl -X GET http://localhost:3000/employee
 [{"id":"61d7f9d64291ec8f94b2ee43","name":"rupert","salary":3424,"age":24},{"id":"63w7f34d654291esdf94b2oa43","name":"rupert","salary":3424,"age":29}]
 # Update the information of the specific employee
 dev@dev:~/go/src/github.com/development/go-fiber-mongo-hrms$ curl -X PUT -d '{"name": "peter", "salary": 3424, "age": 24}' -H 'Content-Type: application/json' http://localhost:3000/employee/61d7f9d64291ec8f94b2ee43
 {"id":"61d7f9d64291ec8f94b2ee43","name":"peter","salary":3424,"age":24}
 # Delete specific employee
  dev@dev:~/go/src/github.com/development/go-fiber-mongo-hrms$ curl -X DELETE http://localhost:3000/employee/61d7f9d64291ec8f94b2ee43
 "record deleted"
```