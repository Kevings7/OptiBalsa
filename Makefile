.PHONY: install test check

install:
	go mod download

check:
	go build ./...
