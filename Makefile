db:
	docker network prune && docker network create grey_db && docker-compose down && docker-compose build --no-cache && docker-compose up

run: 
	go run .

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

.PHONY: db run sqlc test migrateup migratedown buf