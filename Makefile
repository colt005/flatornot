run: build
	@./bin/app

build:
	@CGO_ENABLED=0 go build -o bin/app .