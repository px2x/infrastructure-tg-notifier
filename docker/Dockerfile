# Initial stage: download modules
FROM golang:1.18 as modules

ADD ./go.mod go.sum /modules/
RUN cd /modules && go mod download

# Intermediate stage: Build the binary
FROM golang:1.18 as builder

COPY --from=modules /go/pkg /go/pkg

RUN mkdir -p /service
ADD . /service
WORKDIR /service

# Build the binary with go build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/app ./cmd/main.go

# Final stage: Run the binary
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /service/bin/app /app/service
ADD ./config/config.yaml config/config.yaml
ENTRYPOINT []
CMD ["./service"]
