run:templ build
	@./bin/main
build:
	@go build -o bin/main cmd/main.go
templ:
	@templ generate	