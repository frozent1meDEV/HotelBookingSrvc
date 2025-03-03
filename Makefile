fmt:
	@go fmt ./...

build: fmt
	@go build -o bin/api

run: build
	@./bin/api

seed:
	@go run scripts/seed.go

tests:
	@go test -v ./...

