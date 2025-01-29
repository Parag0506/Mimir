.PHONY: build run test docker

build:
	go build -o api-gateway main.go

run:
	go run main.go

test:
	go test ./tests/...

docker:
	docker build -t mimir-ai-api-gateway . 