build:

	@go build -o bin/shop cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/shop