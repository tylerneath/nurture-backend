GOOSE_DRIVER=postgres
GOOSE_DBSTRING="user=postgres dbname=nuture sslmode=disable password=postgres host=localhost"

.PHONY: migrate

migrate-up:
	goose -dir db/migrations $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

migrate-down:
	goose -dir db/migrations $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down
