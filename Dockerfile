ARG GO_VERSION=1.12

FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir /user && \
    echo 'fbexporter:x:1000:1000:fbexporter:/:' > /user/passwd && \
    echo 'fbexporter:x:1000:' > /user/group

RUN apk add --no-cache ca-certificates git
ENV CGO_ENABLED=0
WORKDIR /src
COPY ./ ./

RUN go build -ldflags '-w -extldflags "-static"'

FROM scratch AS final

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /src/prometheus-flashblade-exporter /usr/bin/prometheus-flashblade-exporter

EXPOSE 9130

USER fbexporter:fbexporter

# Run the compiled binary.
ENTRYPOINT ["/usr/bin/prometheus-flashblade-exporter"]

