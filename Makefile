.PHONY: install test check

install:
	go mod download

test:
	@packages="$$(go list ./...)" || exit 1; \
	if [ -z "$$packages" ]; then \
		echo "No hay paquetes Go todavía"; \
	else \
		go test ./...; \
	fi

check:
	go build ./...