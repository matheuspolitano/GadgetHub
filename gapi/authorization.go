package gapi

import (
	"context"
	"errors"
	"strings"

	"github.com/matheuspolitano/GadgetHub/pkg/tokenManager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) authorizerUser(ctx context.Context, accessibleRole []string) (*tokenManager.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, errors.New("missing authorization header")
	}

	addHeader := values[0]
	fields := strings.Fields(addHeader)
	if len(fields) < 2 {
		return nil, errors.New("not valid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, errors.New("not support authorization type")
	}

	accessToken := fields[1]
	payload, err := server.tokenManager.CheckToken(accessToken)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if !hasPermission(payload.Role, accessibleRole) {
		return nil, errors.New("permission denied")
	}
	return payload, nil

}

func hasPermission(userRole string, accessibleRole []string) bool {
	for _, role := range accessibleRole {
		if userRole == role {
			return true
		}
	}
	return false
}
func unauthenticatedError(err error) error {
	return status.Errorf(codes.Unauthenticated, "unauthorized: %s", err)
}
