# Freecodecamp
Please make sure there's a GO lang installed in your machine. If you're not familiar with GO lang, please go over to their [page](https://go.dev/doc/tutorial/getting-started) or you could visit [**Tour of Go**](https://go.dev/tour/welcome/1). To **download and install**, please follow the instructions describe [here](https://go.dev/doc/install).

Prerequisite:
- GO
- GIT
- GCC Compiler

## Linux installation
```bash
dev@dev:~$ sudo apt install software-properties-common apt-transport-https wget
dev@dev:~$ wget -c https://go.dev/dl/go1.xx.xx.linux-amd64.tar.gz
dev@dev:~$ sudo tar -C /usr/local -xzf go1.xx.xx.linux-amd64.tar.gz
dev@dev:~$ sudo vim .profile
```

Add this inside the .profile
```bash
## GO configuration ##
PATH="$HOME/bin:$HOME/.local/bin:$PATH"
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

**NOTE**: The changes made to the `.profile` will not be applied until the next time you log into your machine. Please restart your machine.

To check if it's already installed:
```bash
dev@dev:~$ go version
go version go1.xx.xx linux/amd64
```

#### GIT Installation
```bash
dev@dev:~$ sudo apt update
dev@dev:~$ sudo apt-get install git
# Git configuration
dev@dev:~$ git config --global user.name "username"
dev@dev:~$ git config --global user.email "email@email.com"
```

#### GCC Installation
```bash
dev@dev:~$ sudo apt update
dev@dev:~$ sudo apt install build-essential
dev@dev:~$ gcc --version
gcc (Ubuntu 7.4.0-1ubuntu1~18.04) 7.4.0
Copyright (C) 2017 Free Software Foundation, Inc.
This is free software; see the source for copying conditions. There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
```

# Table of Contents
### Go Lang
1. [go-server](/go-server/)
2. [go-movies-crud](/go-movies-crud/)
3. [go-bookstore](/go-bookstore/)
4. [slack-age-bot](/slack-bot-age/)
5. [email-checker-tool](/email-checker-tool/)
6. [lambda-basics](/lambda-basics/)
7. [go-fiber-crm-basic](/go-fiber-crm-basic/)
8. [go-fiber-mongo-hrms](/go-fiber-mongo-hrms/)


# Reference
1. [Learn Go Programming by Building 11 Projects](https://www.freecodecamp.org/news/learn-go-by-building-11-projects/)