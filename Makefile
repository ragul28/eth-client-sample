build:
	GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags="-s -w"

install:
	go install
	
run:
	go build && ./eth-client-sample

mod_init:
	go mod init github.com/ragulh28/eth-client-sample
	go get -u

mod:
	go mod tidy
	go mod verify
	go mod vendor