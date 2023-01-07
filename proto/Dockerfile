FROM golang:1.19.3

WORKDIR /app/proto

RUN apt-get update && apt-get install -y protobuf-compiler

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.2

CMD ["make", "generate"]
