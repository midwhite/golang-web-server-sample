FROM golang:1.19.3

WORKDIR /app

RUN apt-get update && apt-get install -y postgresql-client
RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
