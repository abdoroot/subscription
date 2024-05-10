run:templ build
	@./bin/cmd
build:
	@go build -o bin/ ./...
templ:
	@templ generate	
migrate:
	@go run scripts/migrate.go -cmd=up
migrate_down:
	@go run scripts/migrate.go -cmd=down				
test:
	@go test ./... -count=1 -v