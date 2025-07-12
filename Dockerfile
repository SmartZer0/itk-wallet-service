FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o wallet-service ./cmd/app

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/wallet-service .

EXPOSE 8080

CMD ["./wallet-service"]
