.PHONY: install check test

install:
	go mod download

check:
	gofmt -e -l . >/dev/null

test:
	go test ./...
