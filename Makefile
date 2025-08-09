.PHONY: fmt lint tidy check

fmt:
	gofumpt -w .

lint:
	golangci-lint run --timeout=3m

tidy:
	go mod tidy

check: fmt tidy lint
	@echo "OK"
