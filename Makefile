# Makefile for Go Task Manager Project

.PHONY: run test deps tidy lint fmt clean build logs

run:
	go run ./cmd/server/main.go

test:
	go test ./...

deps:
	go mod download

tidy:
	go mod tidy

lint:
	golangci-lint run || echo "Install golangci-lint if needed: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"

fmt:
	gofmt -s -w .

clean:
	rm -rf bin/ coverage.out

build:
	echo "No build step needed for local dev."

logs:
	echo "No logs command for local dev."

# Add more commands as needed
