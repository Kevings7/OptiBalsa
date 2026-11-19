.PHONY: install test check

install:
	go mod download

test: 
	go test ./...

check:
	go build ./...
