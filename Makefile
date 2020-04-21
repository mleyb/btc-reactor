build:
	go build -o bin/btc-reactor main.go

run:
	go run main.go

fmt:
	go fmt ./...

build-image:
	docker build -t btc-reactor:latest . -f Dockerfile