package gapi

import (
	"context"
	"errors"

	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"github.com/matheuspolitano/GadgetHub/pkg/tokenManager"
	"github.com/matheuspolitano/GadgetHub/utils"
	"github.com/matheuspolitano/GadgetHub/val"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	validations := loginValidation(req)
	if len(validations) > 0 {
		return nil, invalidArgumentError(validations)
	}

	user, err := s.store.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	err = utils.CheckPassword(req.GetPassword(), user.HashPassword)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect password")
	}

	accessToken, accessPayload, err := s.tokenManager.GenerateToken(tokenManager.PayloadParameters{
		UserID: int(user.UserID),
		Role:   user.UserRole,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "Error in generated code")
	}

	return &pb.LoginResponse{
		AcessToken:           accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
	}, nil
}

func loginValidation(req *pb.LoginRequest) (validations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		validations = append(validations, fieldViolation("email", err))
	}
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		validations = append(validations, fieldViolation("password", err))
	}
	return
}
