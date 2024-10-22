include .env
export

run:
	go run ./cmd/main.go

database_up:
	docker run --name authorization-dev-db --rm \
	-e POSTGRES_USER=${DB_USER} \
	-e POSTGRES_PASSWORD=${DB_PASS} \
	-e POSTGRES_DB=${DB_NAME} \
	-e PGDATA=/var/lib/postgresql/data \
	-p 5432:5432 \
	-v ./migrations:/docker-entrypoint-initdb.d \
	-d postgres:15.3-bullseye

	clear

database_down:
	docker stop authorization-dev-db

proto_compile:
	protoc --go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			proto/authorization.proto