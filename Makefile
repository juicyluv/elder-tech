.PHONY: openapi
openapi: openapi_http

.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh

.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: mcreate
mcreate:
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: mdrop
mdrop:
	go run cmd/migrate/drop.go

.PHONY: sqlc
sqlc:
	PGPASSWORD=postgres psql -U postgres -d elder -h localhost -f query.sql
