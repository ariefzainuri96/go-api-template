## air-verse for hot reload

1. https://github.com/air-verse/air

## Migration

1. Create migration file:
   migrate create -seq -ext sql -dir ././cmd/migrate/migrations create_users

2. Perform migration:
   migrate -path ./cmd/migrate/migrations -database="postgres://postgres:Rohanwebid96dong@localhost:5433/go-crash-course?sslmode=disable" up

## Docker

1. running docker container that we specify in docker-compose.yml:
   docker compose up -d

2. stop docker container:
   docker compose down

3. remove docker container with its volumes:
   docker compose down -v

## .air.toml

1. current working is for linux because we are using docker for running this apps

2. if you running locally, change .air.toml line 7-8 to:
   bin = "./bin/api.exe"
   cmd = "go build -o ./bin/ ./cmd/api/"

## db connection on host / local machine

1. use address localhost:5433 for connecting to db
