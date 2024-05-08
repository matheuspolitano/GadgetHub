package gapi

import (
	"context"

	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"github.com/matheuspolitano/GadgetHub/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(context context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUserRequest(request)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateName(req.GetFirstName()); err != nil {
		violations = append(violations, fieldViolation("first_name", err))
	}

	if err := val.ValidateName(req.GetLastName()); err != nil {
		violations = append(violations, fieldViolation("last_name", err))
	}

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if err := val.ValidateEmail(req.GetPhone()); err != nil {
		violations = append(violations, fieldViolation("phone", err))
	}
	if err := val.ValidateRole(req.GetUserRole()); err != nil {
		violations = append(violations, fieldViolation("user_role", err))
	}
	return violations
}
