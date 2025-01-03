FROM golang:1.23 AS builder
WORKDIR /iot-controller
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o /iot-controller/cmd/main ./cmd

FROM debian:bookworm-slim
WORKDIR /iot-controller
COPY --from=builder /iot-controller/cmd/main ./cmd/main
COPY .env ./cmd/.env
EXPOSE 50051
CMD ["./cmd/main"]
