#!/bin/bash
set -e

oapi-codegen -generate types -o "./internal/handlers/openapi_types.gen.go" -package "handlers" "./api/app/http.yml"
oapi-codegen -generate chi-server -o "./internal/handlers/openapi_api.gen.go" -package "handlers" "./api/app/http.yml"
