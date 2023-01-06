FROM golang:1.19.3

WORKDIR /app/graphql-gateway

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/cosmtrek/air@latest
