package gapi

import (
	"context"

	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
