start:
	docker start postgres redis

server:
	@/snap/bin/go run cmd/main.go

redis:
	docker run --name redis -p 6379:6379 -d redis
	
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

.PHONY: postgres server redis start