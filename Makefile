# name app
APP_NAME = server

run:
    go run ./cmd/${APP_NAME}/

test:
    go test -coverprofile=coverage.out ./...

coverage:
    go tool cover -html=coverage.out -o coverage.cls
	