.PHONY: run check

run:
	go clean
	go build -o bin/main cmd/main.go
	./bin/main

check:
	go fmt ./...
	go vet ./...