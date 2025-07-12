BINARY_NAME=shell-packet

build:
	go build -o $(BINARY_NAME)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows.exe

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-mac

all: build-linux build-windows build-mac
