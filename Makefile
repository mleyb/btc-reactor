build:
	go build -o bin/btc-reactor main.go

build-lambda:
	GOOS=linux GOARCH=amd64 go build -o bin/btc-reactor main.go
	zip bin/btc-reactor.zip bin/btc-reactor

build-image:
	docker build -t btc-reactor:latest . -f Dockerfile

run:
	go run main.go

fmt:
	go fmt ./...