package gapi

import (
	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func parserUser(user db.User) *pb.User {
	return &pb.User{
		UserId:    int64(user.UserID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		UserRole:  user.UserRole,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
