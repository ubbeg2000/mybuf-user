package main

import (
	"context"
	"fmt"
	"log"
	"net"

	userv1 "github.com/ubbeg2000/mybuf/gen/go/user/v1"
	"google.golang.org/grpc"
)

var users map[string]*userv1.User = make(map[string]*userv1.User)

type UserServiceServer struct {
	userv1.UnimplementedUserServiceServer
}

var _ userv1.UserServiceServer = (*UserServiceServer)(nil)

func (UserServiceServer) AddUser(ctx context.Context, req *userv1.UserServiceAddUserRequest) (*userv1.UserServiceAddUserResponse, error) {
	count := len(users)
	users[fmt.Sprintf("%d", count+1)] = &userv1.User{
		Id:          fmt.Sprintf("%d", count+1),
		Name:        req.Name,
		DateOfBirth: req.DateOfBirth,
	}
	return &userv1.UserServiceAddUserResponse{
		User: users[fmt.Sprintf("%d", count+1)],
	}, nil
}

func (UserServiceServer) GetUsers(context.Context, *userv1.UserServiceGetUsersRequest) (*userv1.UserServiceGetUsersResponse, error) {
	usersSlice := []*userv1.User{}
	for _, v := range users {
		usersSlice = append(usersSlice, v)
	}
	return &userv1.UserServiceGetUsersResponse{
		Users: usersSlice,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	userv1.RegisterUserServiceServer(grpcServer, UserServiceServer{})
	grpcServer.Serve(lis)
}
