package impl

import (
	"context"

	"github.com/midwhite/golang-web-server-sample/grpc-user-service/pb"
)

type UserServiceServer struct{}

func (s *UserServiceServer) GetUserDetail(_ context.Context, params *pb.GetUserDetailParams) (*pb.UserDetail, error) {
	user := pb.UserDetail{
		Id:   params.Id,
		Name: "midwhite",
		Age:  30,
	}
	return &user, nil
}
