.PHONY: check m-up m-down m-sync m-create m-collections pb-serve

all: check

check:
	go fmt ./...
	go vet ./...

m-up: check
	go run cmd/pocketbase/main.go migrate up

m-down:
	go run cmd/pocketbase/main.go migrate down

m-sync:
	go run cmd/pocketbase/main.go migrate history-sync

m-create:
	go run cmd/pocketbase/main.go migrate create temp

m-collections:
	go run cmd/pocketbase/main.go migrate collections

pb-serve: check
	go run cmd/pocketbase/main.go serve