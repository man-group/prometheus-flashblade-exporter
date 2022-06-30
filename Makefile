all: build
build:
	CGO_ENABLED=0 go build -ldflags '-w -extldflags "-static"'
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -extldflags "-static"' -o prometheus-flashblade-exporter.linux
docker:
	docker build
