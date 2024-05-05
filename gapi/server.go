package gapi

import (
	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"github.com/matheuspolitano/GadgetHub/pkg/tokenManager"
)

type Server struct {
	pb.UnimplementedGadgetHubServer
	store        db.Store
	tokenManager tokenManager.Manager
}

func NewServer() Server {
	return Server{}
}
