clean:
	rm greeter

build-go:
	CGO_ENABLED=0 GOOS=linux go build

build-docker:
	docker build -t greeter .

build:
	make clean build-go build-docker