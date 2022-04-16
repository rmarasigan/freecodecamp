# go-movies-crud

![Movies CRUD](/assets/img/go-movies-crud.png "go-movies-crud")

```bash
dev@dev:~$ cd go/src/github.com/development
dev@dev:~/go/src/github.com/development$ mkdir go-movies-crud && cd go-movies-crud
dev@dev:~/go/src/github.com/development/go-movies-crud$ go mod init
go: creating new go.mod: module github.com/development/go-movies-crud
```

## Test API
```bash
dev@dev:~/go/src/github.com/development/go-movies-crud$ go build
dev@dev:~/go/src/github.com/development/go-movies-crud$ go run main.go
dev@dev:~/go/src/github.com/development/go-movies-crud$ ccurl -X POST -d '{"isbn":"4434527","title":"Movie Seven","director":{"firstname":"akhil","lastname":"sharma"}}' -H 'Content-Type: application/json' http://localhost:8000/movies
dev@dev:~/go/src/github.com/development/go-movies-crud$ curl -X PUT -d '{"isbn":"442424","title":"Movie Seven","director":{"firstname":"john","lastname":"mayer"}}' -H 'Content-Type: application/json' http://localhost:8000/movies/74941318
dev@dev:~/go/src/github.com/development/go-movies-crud$ curl -X DELETE http://localhost:8000/movies/74941318
```