.PHONY: check pb-m-up pb-m-down pb-m-sync pb-m-create pb-m-collections pb-serve m-create m-build m-down

all: check

check:
	go fmt ./...
	go vet ./...

pb-m-up: check
	go run cmd/pocketbase/main.go migrate up

pb-m-down:
	go run cmd/pocketbase/main.go migrate down

pb-m-sync:
	go run cmd/pocketbase/main.go migrate history-sync

pb-m-create:
	go run cmd/pocketbase/main.go migrate create temp

pb-m-collections:
	go run cmd/pocketbase/main.go migrate collections

pb-serve: check
	go run cmd/pocketbase/main.go serve

m-create: check
	supabase migration new temp --debug --workdir ./db

m-build: check
	go build -o bin/migration-local cmd/migration-local/main.go

m-down:
	bin/migration-local down