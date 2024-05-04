package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/matheuspolitano/GadgetHub/utils"
	"github.com/stretchr/testify/require"
)

func generateUserParams() CreateUserParams {
	return CreateUserParams{
		FirstName:    utils.RandString(10),
		LastName:     utils.RandString(10),
		Email:        utils.RandEmail(),
		HashPassword: utils.RandString(20),
		Phone:        utils.RandPhone(),
		IsAdmin:      true,
	}
}
func createUserTest(t *testing.T) User {
	userParms := generateUserParams()
	userCreated, err := testQuerier.CreateUser(context.TODO(), userParms)
	require.NoError(t, err)
	require.NotEmpty(t, userCreated)
	return userCreated
}

func TestCreateUser(t *testing.T) {
	userParms := generateUserParams()
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

func TestGetUser(t *testing.T) {
	userParms := generateUserParams()
	userCreated, err := testQuerier.CreateUser(context.TODO(), userParms)
	require.NoError(t, err)
	require.NotEmpty(t, userCreated)

	user, err := testQuerier.GetUser(context.TODO(), userCreated.UserID)
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

func TestGetUserByEmail(t *testing.T) {
	userParms := generateUserParams()
	userCreated, err := testQuerier.CreateUser(context.TODO(), userParms)
	require.NoError(t, err)
	require.NotEmpty(t, userCreated)

	user, err := testQuerier.GetUserByEmail(context.TODO(), userCreated.Email)
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

func TestGetUserByPhone(t *testing.T) {
	userParms := generateUserParams()
	userCreated, err := testQuerier.CreateUser(context.TODO(), userParms)
	require.NoError(t, err)
	require.NotEmpty(t, userCreated)

	user, err := testQuerier.GetUserByPhone(context.TODO(), userCreated.Phone)
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

func TestUpdateOnlyPassword(t *testing.T) {
	oldUser := createUserTest(t)
	newPassword := utils.RandString(10)

	userUpdated, err := testQuerier.UpdateUser(context.TODO(), UpdateUserParams{
		UserID: oldUser.UserID,
		HashPassword: pgtype.Text{
			String: newPassword,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, userUpdated)

	require.Equal(t, newPassword, userUpdated.HashPassword)

	err = testQuerier.DeleteUser(context.Background(), userUpdated.UserID)
	require.NoError(t, err)
}

func TestUpdateOnlyFirstName(t *testing.T) {
	oldUser := createUserTest(t)
	newFirstName := utils.RandString(15)

	userUpdated, err := testQuerier.UpdateUser(context.TODO(), UpdateUserParams{
		UserID: oldUser.UserID,
		FirstName: pgtype.Text{
			String: newFirstName,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, userUpdated)

	require.Equal(t, newFirstName, userUpdated.FirstName)

	err = testQuerier.DeleteUser(context.Background(), userUpdated.UserID)
	require.NoError(t, err)
}

func TestUpdateOnlyLastName(t *testing.T) {
	oldUser := createUserTest(t)
	newLastName := utils.RandString(15)

	userUpdated, err := testQuerier.UpdateUser(context.TODO(), UpdateUserParams{
		UserID: oldUser.UserID,
		LastName: pgtype.Text{
			String: newLastName,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, userUpdated)

	require.Equal(t, newLastName, userUpdated.LastName)

	err = testQuerier.DeleteUser(context.Background(), userUpdated.UserID)
	require.NoError(t, err)
}

func TestUpdateOnlyEmail(t *testing.T) {
	oldUser := createUserTest(t)
	newEmail := utils.RandEmail()

	userUpdated, err := testQuerier.UpdateUser(context.TODO(), UpdateUserParams{
		UserID: oldUser.UserID,
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, userUpdated)

	require.Equal(t, newEmail, userUpdated.Email)

	err = testQuerier.DeleteUser(context.Background(), userUpdated.UserID)
	require.NoError(t, err)
}

func TestUpdateOnlyPhone(t *testing.T) {
	oldUser := createUserTest(t)
	newPhone := utils.RandPhone()

	userUpdated, err := testQuerier.UpdateUser(context.TODO(), UpdateUserParams{
		UserID: oldUser.UserID,
		Phone: pgtype.Text{
			String: newPhone,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, userUpdated)

	require.Equal(t, newPhone, userUpdated.Phone)

	err = testQuerier.DeleteUser(context.Background(), userUpdated.UserID)
	require.NoError(t, err)
}
