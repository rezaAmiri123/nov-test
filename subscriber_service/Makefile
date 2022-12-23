MIGRATION_PATH=internal/adapters/migration
DB_NAME=microservice
DB_USER=postgres
DB_PASSWORD=postgres
DB_ADDR=nov_postgesql
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_ADDR):5432/$(DB_NAME)?sslmode=disable

migratecreate:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq subscriber_service

psql:
	psql -h $(DB_ADDR) -U $(DB_USER) -d $(DB_NAME) -W

run:
	go run cmd/subscriber/main.go

migrateup:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down 1

migrateforce:
	migrate  -path $(MIGRATION_PATH) -database "$(DB_URL)" force 1

test_coverage:
	go test --cover ./...
