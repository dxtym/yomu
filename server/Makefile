server:
	@/snap/bin/go run cmd/main.go

test:
	@/snap/bin/go test -v ./....

build:
	@/snap/bin/go build -o bin/main cmd/main.go

.PHONY: server test build