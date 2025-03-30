FROM golang:1.24.0-bookworm AS base

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY ./database/migrations/ ./database/migrations/

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["air"]
