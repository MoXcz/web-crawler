run: build
	@./crawler https://mocosmo.me 5 10

build:
	@go build -o crawler

test:
	@go test -v ./...
