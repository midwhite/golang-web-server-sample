package impl

import (
	"context"

	"github.com/midwhite/golang-web-server-sample/grpc-user-service/pb"
)

type UserServiceServer struct{}

func (s *UserServiceServer) GetUsers(_ context.Context, _ *pb.GetUsersParams) (*pb.UserList, error) {
	users := []*pb.User{
		{Id: "1", Name: "北条氏康", Age: 20},
		{Id: "2", Name: "武田信玄", Age: 30},
		{Id: "3", Name: "今川義元", Age: 40},
	}
	response := pb.UserList{
		Users: users,
	}
	return &response, nil
}

func (s *UserServiceServer) CreateUser(_ context.Context, _ *pb.CreateUserParams) (*pb.User, error) {
	user := pb.User{
		Id:   "uuid",
		Name: "長尾景虎",
		Age:  25,
	}
	return &user, nil
}

func (s *UserServiceServer) UpdateUser(_ context.Context, params *pb.UpdateUserParams) (*pb.User, error) {
	user := pb.User{
		Id:   params.Id,
		Name: "上杉謙信",
		Age:  35,
	}
	return &user, nil
}
