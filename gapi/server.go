package gapi

import (
	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"github.com/matheuspolitano/GadgetHub/pkg/tokenManager"
	"github.com/matheuspolitano/GadgetHub/utils"
)

type Server struct {
	pb.UnimplementedGadgetHubServer
	store        db.Store
	tokenManager tokenManager.Manager
	config       utils.Config
}

func NewServer(store db.Store, config utils.Config) (*Server, error) {
	tokenMnager, err := tokenManager.NewJWTManager(config.TokenSecretKey, config.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	return &Server{
		store:        store,
		tokenManager: tokenMnager,
		config:       config,
	}, nil
}
