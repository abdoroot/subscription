export PATH := $(PATH):$(HOME)/go/bin
run:export templ build
	@./bin/cmd
build:
	@go build -o bin/ ./...
templ:
	@templ generate	
migrate:createdb
	@go run scripts/migrate.go -cmd=up
migrate_down:
	@go run scripts/migrate.go -cmd=down				
test:
	@go test ./... -count=1 -v
css:
	npx tailwindcss -i views/css/app.css -o static/css/style.css --watch
watch:export
	templ generate --watch --proxy="http://localhost:4000" --cmd="air air -c .air.tom"
createdb:
	@docker exec -it pq-container createdb --username=postgres --owner=postgres subscription
dropdb:
	@docker exec -it pq-container dropdb --username=postgres subscription		
.PHONY: export