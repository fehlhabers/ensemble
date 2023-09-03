build:
	@go build -o bin/ensemble .

test:
	@go test -v ./...
