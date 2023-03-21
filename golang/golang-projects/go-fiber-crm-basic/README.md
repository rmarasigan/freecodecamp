# go-fiber-crm-basic

Prerequisite:
- SQLite installed

## Install SQLiteBrowser
```bash
dev@dev:~:$ sudo apt update
dev@dev:~:$ sudo apt install sqlitebrowser
dev@dev:~:$ sqlitebrowser --version
This is DB Browser for SQLite version 3.11.2.
# Uninstall or Remove SQLite Browser
dev@dev:~:$ sudo apt --purge remove sqlitebrowser
```

## Development
```bash
dev@dev:~$ cd go/src/github.com/development
dev@dev:~/go/src/github.com/development$ mkdir go-fiber-crm-basic && cd go-fiber-crm-basic
dev@dev:~/go/src/github.com/development/go-fiber-crm-basic$ go mod init
go: creating new go.mod: module github.com/development/go-fiber-crm-basic
dev@dev:~/go/src/github.com/development/go-fiber-crm-basic$ go get "github.com/gofiber/fiber"
go: downloading github.com/gofiber/fiber v1.14.6
go: downloading github.com/gofiber/utils v0.0.10
go: downloading github.com/gorilla/schema v1.1.0
go: downloading github.com/mattn/go-colorable v0.1.7
go: downloading github.com/mattn/go-isatty v0.0.12
go: downloading github.com/valyala/fasthttp v1.16.0
go: downloading golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980
go: downloading github.com/valyala/tcplisten v0.0.0-20161114210144-ceec8f93295a
go: downloading github.com/klauspost/compress v1.10.7
go: downloading github.com/andybalholm/brotli v1.0.0
go get: added github.com/andybalholm/brotli v1.0.0
go get: added github.com/gofiber/fiber v1.14.6
go get: added github.com/gofiber/utils v0.0.10
go get: added github.com/gorilla/schema v1.1.0
go get: added github.com/klauspost/compress v1.10.7
go get: added github.com/mattn/go-colorable v0.1.7
go get: added github.com/mattn/go-isatty v0.0.12
go get: added github.com/valyala/bytebufferpool v1.0.0
go get: added github.com/valyala/fasthttp v1.16.0
go get: added github.com/valyala/tcplisten v0.0.0-20161114210144-ceec8f93295a
go get: added golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980
# It will generate an executable file, named main
dev@dev:~/go/src/github.com/development/go-fiber-crm-basic$ go build main.go
# Run your executable file. After running, you should see the `Connection opened to database`.
dev@dev:~/go/src/github.com/development/go-fiber-crm-basic$ go run main.go
Connection opened to database
Database Migrated
        _______ __                 
  ____ / ____(_) /_  ___  _____   HOST     0.0.0.0  OS      LINUX
_____ / /_  / / __ \/ _ \/ ___/   PORT     3000     THREADS 8
  __ / __/ / / /_/ /  __/ /       TLS      FALSE    PREFORK FALSE
    /_/   /_/_.___/\___/_/1.14.6  HANDLERS 4        PID     108735
v2 will be released on 15 September 2020!
Please visit https://gofiber.io/v2 for more information.
```

## Test API
```bash
# Creating Lead
dev@dev:~/go/src/github.com/development/go-fiber-crm-basic$ curl -X POST -d '{"name": "Ramsey Peters", "company": "Amazon", "email": "alex@amazon.com", "phone": 123456789}' -H 'Content-Type: application/json' http://localhost:3000/api/v1/lead
{"ID":1,"CreatedAt":"2022-04-16T16:07:47.302337828+08:00","UpdatedAt":"2022-04-16T16:07:47.302337828+08:00","DeletedAt":null,"name":"Ramsey Peters","company":"Amazon","email":"alex@amazon.com","phone":123456789}
dev@dev:~/go/src/github.com/development/go-fiber-crm-basic$ curl -X POST -d '{"name": "Alex Clinton", "company": "Microsoft", "email": "alex@microsoft.com", "phone": 987654321}' -H 'Content-Type: application/json' http://localhost:3000/api/v1/lead
{"ID":2,"CreatedAt":"2022-04-16T16:09:31.283965903+08:00","UpdatedAt":"2022-04-16T16:09:31.283965903+08:00","DeletedAt":null,"name":"Alex Clinton","company":"Microsoft","email":"alex@microsoft.com","phone":987654321}
# Getting all Leads
dev@dev:~/go/src/github.com/development/go-fiber-crm-basic$ curl -X GET http://localhost:3000/api/v1/lead
[{"ID":1,"CreatedAt":"2022-04-16T16:07:47.302337828+08:00","UpdatedAt":"2022-04-16T16:07:47.302337828+08:00","DeletedAt":null,"name":"Ramsey Peters","company":"Amazon","email":"alex@amazon.com","phone":123456789},{"ID":2,"CreatedAt":"2022-04-16T16:09:31.283965903+08:00","UpdatedAt":"2022-04-16T16:09:31.283965903+08:00","DeletedAt":null,"name":"Alex Clinton","company":"Microsoft","email":"alex@microsoft.com","phone":987654321}]
# Getting specific Lead
dev@dev:~/go/src/github.com/development/go-fiber-crm-basic$ curl -X GET http://localhost:3000/api/v1/lead/1
{"ID":1,"CreatedAt":"2022-04-16T16:07:47.302337828+08:00","UpdatedAt":"2022-04-16T16:07:47.302337828+08:00","DeletedAt":null,"name":"Ramsey Peters","company":"Amazon","email":"alex@amazon.com","phone":123456789}
# Deleting specific Lead data
dev@dev:~/go/src/github.com/development/go-fiber-crm-basic$ curl -X DELETE http://localhost:3000/api/v1/lead/1
Lead successfully deleted
dev@dev:~/go/src/github.com/development/go-fiber-crm-basic$ curl -X GET http://localhost:3000/api/v1/lead
[{"ID":2,"CreatedAt":"2022-04-16T16:09:31.283965903+08:00","UpdatedAt":"2022-04-16T16:09:31.283965903+08:00","DeletedAt":null,"name":"Alex Clinton","company":"Microsoft","email":"alex@microsoft.com","phone":987654321}]
```