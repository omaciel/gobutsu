createdb:
	docker exec -it postgres createdb --owner=root --encoding=UTF8 --lc-collate=en_US.utf8 --lc-ctype=en_US.utf8 --tablespace=pg_default gobutsu

dropdb:
	docker exec -it postgres dropdb gobutsu

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gobutsu?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gobutsu?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gobutsu?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gobutsu?sslmode=disable" -verbose down 1

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

test:
	go test -v -cover ./...

server:
	go run main.go

sqlc:
	sqlc generate

.PHONY: createdb dropdb migratedown migratedown1 migrateup migrateup1 postgres server sqlc test