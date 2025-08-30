
run:
	air

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

dotenv-migrate-up:
	set -a; source .env; set +a; migrate -path ./migrations -database "$$DATABASE_URL" up

dotenv-migrate-down:
	set -a; source .env; set +a; migrate -path ./migrations -database "$$DATABASE_URL" down 1
