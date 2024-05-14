package gapi

import (
	"context"

	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"github.com/matheuspolitano/GadgetHub/utils"
	"github.com/matheuspolitano/GadgetHub/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, err := s.authorizerUser(ctx, []string{"admin"})
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateCreateUserRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}
	hash_password, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}
	userParams := db.CreateUserParams{
		FirstName:    req.GetFirstName(),
		LastName:     req.GetLastName(),
		Email:        req.GetEmail(),
		Phone:        req.GetPhone(),
		HashPassword: hash_password,
		UserRole:     req.GetUserRole(),
	}

	createResult, err := s.store.CreateUser(ctx, userParams)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user %s", err)
	}

	return &pb.CreateUserResponse{
		User: parserUser(createResult),
	}, nil
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

	if err := val.ValidatePhone(req.GetPhone()); err != nil {
		violations = append(violations, fieldViolation("phone", err))
	}
	if err := val.ValidateRole(req.GetUserRole()); err != nil {
		violations = append(violations, fieldViolation("user_role", err))
	}
	return violations
}
