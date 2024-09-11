all:
	@go run cmd/main.go

test:
	go test -v ./cmd/
