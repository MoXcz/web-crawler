run: build
	@./crawler https://crawler-test.com/ 5 10

build:
	@go build -o crawler

test:
	@go test -v ./...
