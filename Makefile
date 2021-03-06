postgrescreate:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

postgres:
	docker start postgres12

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

mysqlcreate:
	docker run --name mysql8 -p 3307:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:latest

mysql:
	docker start mysql8
	
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

server:
	go run main.go

mock:
	mockgen -package mockapi -destination db/mock/handler.go simple_bank/api Handler

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server test mock