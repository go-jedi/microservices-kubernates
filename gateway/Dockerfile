FROM golang:1.24.0-alpine AS builder

WORKDIR /github.com/go-jedi/gateway/app
COPY . /github.com/go-jedi/gateway/app

RUN go mod download
RUN go build -ldflags="-s -w" -trimpath -buildvcs=false -o .bin/app cmd/app/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/go-jedi/gateway/app/.bin/app .
COPY config.yaml /root/

CMD ["./app", "--config", "config.yaml"]