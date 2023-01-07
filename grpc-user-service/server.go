package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/midwhite/golang-web-server-sample/grpc-user-service/impl"
	"github.com/midwhite/golang-web-server-sample/grpc-user-service/pb"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("grpc-user-service:%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, &impl.UserServiceServer{})
	grpcServer.Serve(lis)
}
