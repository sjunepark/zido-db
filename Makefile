.PHONY: run check

run-vector:
	go clean
	go build -o bin/vector cmd/vector/main.go
	./bin/vector

run-location:
	go clean
	go build -o bin/location cmd/location/main.go
	./bin/location

check:
	go fmt ./...
	go vet ./...