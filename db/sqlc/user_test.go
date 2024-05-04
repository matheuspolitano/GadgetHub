package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	userParms := CreateUserParams{
		FirstName:    "Matheus",
		LastName:     "Politano",
		Email:        "Matheus Politano",
		HashPassword: "2139u23jeio93eu12",
		Phone:        "+12345678",
		IsAdmin:      true,
	}
	user, err := testQuerier.CreateUser(context.TODO(), userParms)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, userParms.FirstName, user.FirstName)
	require.Equal(t, userParms.LastName, user.LastName)
	require.Equal(t, userParms.Email, user.Email)
	require.Equal(t, userParms.HashPassword, user.HashPassword)
	require.Equal(t, userParms.IsAdmin, user.IsAdmin)

	require.NotZero(t, user.UserID)
	err = testQuerier.DeleteUser(context.Background(), user.UserID)
	require.NoError(t, err)
}
