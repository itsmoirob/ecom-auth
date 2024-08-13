build:
	@go build -o bin/ecom-auth cmd/main.go

run: build
	@./bin/ecom-auth

test:	
	@go test -v ./...

migration:
	@migrate create -ext sql -dir /home/robbie/source/go/ecom-auth/cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down