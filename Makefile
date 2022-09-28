# Commands related to the database
db-start:
	docker-compose -f docker-compose.local.yml up -d

db-stop:
	docker-compose -f docker-compose.local.yml stop

db-kill:
	docker-compose -f docker-compose.local.yml down -v

# Update database schema after an update
schema-generate:
	go generate ./ent

schema-view:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

# Start server
go-server:
	go run api/main.go

go-test:
	go test ./...

go-format: # Run go format to format files
	go fmt ./...

go-build:
	go build -o app api/*.go