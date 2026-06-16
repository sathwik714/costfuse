.PHONY: build run test lint clean

build:
	go build -o bin/costfuse ./cmd/costfuse

run:
	go run ./cmd/costfuse --dry-run

test:
	go test ./...

lint:
	go vet ./...

clean:
	rm -rf bin
