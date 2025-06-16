.PHONY: run swag test seed migrate reset seeddata

run:
	go run cmd/main.go

swag:
	swag init --generalInfo cmd/main.go --output docs

test:
	go test ./...

migrate:
	bash -c 'set -a && source .env && psql -d "$$DATABASE_URL" -f database/migration.sql'

reset:
	bash -c 'set -a && source .env && psql -d "$$DATABASE_URL" -f database/reset.sql'

seeddata:
	bash -c 'set -a && source .env && psql -d "$$DATABASE_URL" -f database/seed.sql'

seed: reset migrate seeddata
