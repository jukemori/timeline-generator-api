.PHONY: seed gqlgen run

seed:
	@go run cmd/seed/main.go

gqlgen:
	@go run github.com/99designs/gqlgen generate

run:
	@go run cmd/server/main.go 