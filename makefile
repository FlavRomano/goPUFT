include .env

build:
	go build -o ./bin/client ./cmd/client/main.go
	go build -o ./bin/server ./cmd/server/main.go

.PHONY: clean
clean:
	rm ./bin/*