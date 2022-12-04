go generate ./...
run postgres
docker run --env-file ./.env -p 5432:5432 postgres
remember to run CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
go run main.go