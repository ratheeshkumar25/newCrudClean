run:
		go run main.go

test:
		go test -v ./...
deps:
		go mod tidy

std:
		golint ./...