# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
COPY data.json .

RUN CGO_ENABLED=0 GOOS=linux go build -o /item-comparison-ai-api ./cmd/api

# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /item-comparison-ai-api .
COPY --from=builder /app/data.json .

ENV DATA_FILE_PATH=/root/data.json
ENV BIND_ADDR=:8080
EXPOSE 8080

CMD ["./item-comparison-ai-api"]