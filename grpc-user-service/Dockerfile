FROM golang:1.19.3

WORKDIR /app/grpc-user-service

RUN apt-get update && apt-get install -y postgresql-client

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/cosmtrek/air@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.1.0
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

CMD ["make", "start"]
