package server

import (
	"context"

	"github.com/midwhite/golang-web-server-sample/grpc-user-service/pb/github.com/midwhite/golang-web-serber-sample/grpc-user-service/pb"
)

type UserServiceServer struct{}

func (s *UserServiceServer) GetUserDetail(context.Context, *pb.GetUserDetailParams) (*pb.UserDetail, error) {
	return nil, nil
}
