package userservice

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/midwhite/golang-web-server-sample/graphql-gateway/pb"
)

var Client pb.UserServiceClient

func Setup() (func(), error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("grpc-user-service:50051", opts...)
	if err != nil {
		return nil, err
	}
	Client = pb.NewUserServiceClient(conn)

	return func() {
		conn.Close()
	}, nil
}
