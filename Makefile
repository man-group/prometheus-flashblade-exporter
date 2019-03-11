all: build
build:
	CGO_ENABLED=0 go build -ldflags '-w -extldflags "-static"'
docker:
	docker build
