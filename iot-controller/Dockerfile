FROM golang:1.23 AS builder

WORKDIR /iot-controller
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /iot-controller/main ./cmd

FROM debian:bookworm-slim
WORKDIR /iot-controller
COPY .env ./cmd/.env
COPY --from=builder /iot-controller/main ./main
EXPOSE 50051
CMD ["./main"]
