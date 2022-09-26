DOCKER_REGISTRY                := sourabh181097
DOCKER_REPOSITORY              := darwin
DOCKER_IMG                     := $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY)
TAG 	                       := 0.2.0

.PHONY: go-build
go-build:
	go build -o ./bin/darwin ./src/darwin

.PHONY: build
build:
	docker buildx build --platform linux/amd64 -t $(DOCKER_IMG):$(TAG) .

.PHONY: start
start: build
	docker run --rm --name=darwin -d -p 8080:8080 $(DOCKER_IMG):$(TAG)

.PHONY: stop
stop:
	docker stop darwin

.PHONY: clean
clean:
	docker rmi $(DOCKER_IMG):$(TAG)

.PHONY: push
push:
	docker push $(DOCKER_IMG):$(TAG)
