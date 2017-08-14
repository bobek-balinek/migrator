.PHONY: build

build:
	@go build -o ./bin/migrate cmd/migrate/migrate.go
