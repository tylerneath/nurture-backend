GOOSE_DRIVER=postgres
GOOSE_DBSTRING="user=postgres dbname=nuture sslmode=disable password=postgres host=localhost"
MIGRATION_FILE=db/migrations/20240301190841_drop_tables.sql

.PHONY: migrate

reset-db: drop-tables migrate

migrate:
	go run main.go migrate

drop-tables:
	goose -dir db/migrations ${GOOSE_DRIVER} ${GOOSE_DBSTRING} up ${MIGRATION_FILE}