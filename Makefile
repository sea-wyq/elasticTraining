all: build

build:
	go build -o ./elastic main.go

Host ssh1
	HostName 172.17.0.2
Host ssh2
	HostName 172.17.0.3