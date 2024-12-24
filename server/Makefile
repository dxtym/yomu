
dropdb:
	docker exec -it postgres dropdb yomu

createdb:
	docker exec -it postgres createdb --username=root --owner=root yomu

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

.PHONY: createdb dropdb postgres