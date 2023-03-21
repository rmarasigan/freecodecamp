# go-bookstore

![Go Bookstore Structure](/assets/img/go-bookstore.png "go-bookstore")

```bash
dev@dev:~$ cd go/src/github.com/development
dev@dev:~/go/src/github.com/development$ mkdir go-bookstore && cd go-bookstore
dev@dev:~/go/src/github.com/development/go-bookstore$ go mod init
go: creating new go.mod: module github.com/development/go-bookstore
dev@dev:~/go/src/github.com/development/go-bookstore$ go get "github.com/jinzhu/gorm"
go: downloading github.com/jinzhu/gorm v1.9.16
go: downloading github.com/jinzhu/inflection v1.0.0
go get: added github.com/jinzhu/gorm v1.9.16
go get: added github.com/jinzhu/inflection v1.0.0
dev@dev:~/go/src/github.com/development/go-bookstore$ go get "github.com/gorilla/mux"
go: downloading github.com/gorilla/mux v1.8.0
```

 ## Test API
 ```bash
dev@dev:~/go/src/github.com/development/go-bookstore$ curl -X GET http://localhost:9010/book/
dev@dev:~/go/src/github.com/development/go-bookstore$ curl -X POST -d '{"Name":"Zero to One","Author":"Peter Thiel","Publication":"Penguin"}' -H 'Content-Type: application/json' http://localhost:9010/book/
{"ID":1,"CreatedAt":"2022-04-15T14:17:03.085062718+08:00","UpdatedAt":"2022-04-15T14:17:03.085062718+08:00","DeletedAt":null,"name":"Zero to One","author":"Peter Thiel","publication":"Penguin"}

dev@dev:~/go/src/github.com/development/go-bookstore$ curl -X POST -d '{"Name": "The Startup way", "Author": "Eric Ries", "Publication": "Penguin"}' -H 'Content-Type: application/json' http://localhost:9010/book/
{"ID":2,"CreatedAt":"2022-04-15T14:20:04.32792735+08:00","UpdatedAt":"2022-04-15T14:20:04.32792735+08:00","DeletedAt":null,"name":"The Startup way","author":"Eric Ries","publication":"Penguin"}

dev@dev:~/go/src/github.com/development/go-bookstore$ curl -X GET http://localhost:9010/book/1
{"ID":1,"CreatedAt":"2022-04-15T06:17:03Z","UpdatedAt":"2022-04-15T06:17:03Z","DeletedAt":null,"name":"Zero to One","author":"Peter Thiel","publication":"Penguin"}

dev@dev:~/go/src/github.com/development/go-bookstore$ curl -X GET http://localhost:9010/book/2
{"ID":2,"CreatedAt":"2022-04-15T06:20:04Z","UpdatedAt":"2022-04-15T06:20:04Z","DeletedAt":null,"name":"The Startup way","author":"Eric Ries","publication":"Penguin"}

dev@dev:~/go/src/github.com/development/go-bookstore$ curl -X PUT -d '{"Name": "The Startup way", "Author": "Eric Ries", "Publication": "Orion"}' -H 'Content-Type: application/json' http://localhost:9010/book/2
{"ID":2,"CreatedAt":"2022-04-15T06:20:04Z","UpdatedAt":"2022-04-15T14:32:40.823071058+08:00","DeletedAt":null,"name":"The Startup way","author":"Eric Ries","publication":"Orion"}

dev@dev:~/go/src/github.com/development/go-bookstore$ curl -X DELETE http://localhost:9010/book/2
{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"","author":"","publication":""}