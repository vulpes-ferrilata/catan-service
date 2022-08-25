dns := mysql://catan:123456@tcp(localhost:3306)/catan?charset=utf8mb4
path := ./migration

migrate:
	migrate -database ${dns} -path ${path} up

generate-swag:
	swag init --pd -d cmd/http

.PHONY: migrate generate-swag