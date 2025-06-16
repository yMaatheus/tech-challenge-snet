.PHONY: run swag test migrate

run:
	go run cmd/main.go

swag:
	swag init --generalInfo cmd/main.go --output docs

test:
	go test ./...

migrate:
	psql $(DATABASE_URL) -f database/migration.sql
