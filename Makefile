run: build
	@./crawler https://mocosmo.me

build:
	@go build -o crawler

test:
	@go test -v ./...
