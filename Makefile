.PHONY: check pb-m-up pb-m-down pb-m-sync pb-m-create pb-m-collections pb-serve m-create m-build m-down

all: check

check:
	go fmt ./...
	go vet ./...

m-create: check
	supabase migration new temp --debug --workdir ./db

m-build: check
	go build -o bin/migration-local cmd/migration-local/main.go

m-down:
	bin/migration-local down

m-status:
	bin/migration-local status