.PHONY: run dev swag test coverage seed migrate reset seeddata

run:
	go run cmd/main.go

dev:
	air

swag:
	swag init --generalInfo cmd/main.go --output docs

test:
	bash -c 'set -a && source .env && go test ./...'

coverage:
	bash -c 'set -a && source .env && go test -cover -coverpkg=./repository,./handler,./service ./repository ./handler ./service'

migrate:
	bash -c 'set -a && source .env && psql -d "$$DATABASE_URL" -f database/migration.sql'

reset:
	bash -c 'set -a && source .env && psql -d "$$DATABASE_URL" -f database/reset.sql'

seeddata:
	bash -c 'set -a && source .env && psql -d "$$DATABASE_URL" -f database/seed.sql'

seed: reset migrate seeddata
