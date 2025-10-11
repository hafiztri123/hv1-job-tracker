SHELL := /bin/bash

ifneq (,$(wildcard .env))
	include .env
	export
endif

DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(APP_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable


db-create:
	@migrate create -ext sql -dir ./migrations -seq $(v)

db-up:
	@migrate -database "$(DB_URL)" -path migrations up

db-down:
	@migrate -database "$(DB_URL)" -path migrations down


db-ver:
	@migrate -database "$(DB_URL)" -path migrations version

db-force:
	@migrate -database "$(DB_URL)" -path migrations force $(v)


test:
	@go test -coverprofile=coverage.out ./internal/...
	@go tool cover -func=coverage.out

test-html:
	@go test -coverprofile=coverage.out ./internal/...
	@go tool cover -html=coverage.out
