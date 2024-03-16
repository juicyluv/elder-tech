.PHONY: openapi
openapi: openapi_http

.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh users internal/users main

.PHONY: run
run:
	go run cmd/app/main.go