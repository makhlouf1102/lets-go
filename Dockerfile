FROM golang:1.23.2

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go install github.com/air-verse/air@latest

RUN go mod download

COPY . .