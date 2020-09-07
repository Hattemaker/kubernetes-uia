all: clean build-go build-docker

.PHONY: clean
clean:
	rm greeter

.PHONY: build-go
build-go:
	CGO_ENABLED=0 GOOS=linux go build

.PHONY: build-docker
build-docker:
	docker build -t greeter .
