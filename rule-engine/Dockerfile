FROM golang:1.23 AS builder

WORKDIR /rule-engine
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /rule-engine/main ./cmd

FROM debian:bookworm-slim
WORKDIR /rule-engine
COPY .env ./cmd/.env
COPY --from=builder /rule-engine/main ./main
EXPOSE 50051
CMD ["./main"]
