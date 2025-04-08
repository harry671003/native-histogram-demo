FROM golang:1.24 AS builder

ENV GOPROXY=direct

COPY go.mod .
COPY go.sum .
RUN go mod download -x

COPY . .
RUN go build -o /demo cmd/demo/main.go

FROM debian:bookworm-slim

COPY --from=builder /demo /demo
ENTRYPOINT [ "/demo" ]
