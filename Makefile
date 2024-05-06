run:templ build
	@./bin/cmd
build:
	@go build -o bin/ ./...
templ:
	@templ generate	
migrate:
	@go run scripts/migrate.go		