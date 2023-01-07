FROM golang:1.19.3

WORKDIR /app/graphql-gateway

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/cosmtrek/air@v1.40.4
RUN go install github.com/99designs/gqlgen@v0.17.22
