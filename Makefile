run:
	docker-compose down && docker-compose build --no-cache && docker-compose --env-file ./.env up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

migrateup:
	migrate -path db/migrations -database ${DB_URI} -verbose up

migratedown:
	migrate -path db/migrations -database ${DB_URI} -verbose down

buf:
	buf mod update && buf generate

.PHONY: run sqlc test migrateup migratedown buf