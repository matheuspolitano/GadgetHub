package gapi

import (
	"context"

	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func validateCreateUserReuqest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.FistName(req.GetFirstName()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}
