.PHONY: run migrate-up migrate-down migrate-create

# Carregar as variáveis do arquivo .env
include .env
export

run:
	docker-compose -f .trvium/docker-compose.yml up --build

migrate-up:
	migrate -path ./db/migrations -database "$(DATABASE_URL)" -verbose up

migrate-down:
	migrate -path ./db/migrations -database "$(DATABASE_URL)" -verbose down 1

migrate-create:
	migrate create -ext sql -dir ./db/migrations -seq $(name)