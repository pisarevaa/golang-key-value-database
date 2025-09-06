FROM golang:1.24.2-alpine AS builder

WORKDIR /opt/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o bin/kvdatabase ./cmd

FROM alpine:3.21
RUN addgroup -S kvdatabase && adduser -S kvdatabase -G kvdatabase
USER kvdatabase
WORKDIR /app
COPY .env.example .env
COPY --from=builder /opt/app/bin/kvdatabase /usr/bin/kvdatabase

USER root
# CMD ["kvdatabase"]
