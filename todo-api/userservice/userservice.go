package userservice

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/midwhite/golang-web-server-sample/todo-api/pb"
)

func GetUserDetail(id string) *pb.UserDetail {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("grpc-user-service:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	user, err := client.GetUserDetail(context.Background(), &pb.GetUserDetailParams{Id: id})

	if err != nil {
		log.Fatalf("client.GetUserDetail failed: %v", err)
	}

	return user
}
