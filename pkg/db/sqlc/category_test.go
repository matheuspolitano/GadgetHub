package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/matheuspolitano/GadgetHub/utils"
	"github.com/stretchr/testify/require"
)

func generateCategoryParams() CreateCategoryParams {
	return CreateCategoryParams{
		Name:        utils.RandString(10),
		Description: utils.RandString(50),
	}
}

func createCategoryTest(t *testing.T) Category {
	catParams := generateCategoryParams()
	category, err := testQuerier.CreateCategory(context.TODO(), catParams)
	require.NoError(t, err)
	require.NotEmpty(t, category)
	return category
}
func TestCreateCategory(t *testing.T) {
	catParams := generateCategoryParams()
	category, err := testQuerier.CreateCategory(context.TODO(), catParams)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, catParams.Name, category.Name)
	require.Equal(t, catParams.Description, category.Description)

	// Clean up
	err = testQuerier.DeleteCategory(context.Background(), category.CategoryID)
	require.NoError(t, err)
}

func TestGetCategoriesByName(t *testing.T) {
	category := createCategoryTest(t)

	fetchedCategories, err := testQuerier.GetCategoriesByName(context.TODO(), "%"+category.Name+"%")
	require.NoError(t, err)
	require.NotEmpty(t, fetchedCategories)
	require.True(t, len(fetchedCategories) > 0)

	// Clean up
	err = testQuerier.DeleteCategory(context.Background(), category.CategoryID)
	require.NoError(t, err)
}

func TestUpdateCategory(t *testing.T) {
	originalCategory := createCategoryTest(t)
	newName := utils.RandString(10)
	newDescription := utils.RandString(50)

	updatedCategory, err := testQuerier.UpdateCategory(context.TODO(), UpdateCategoryParams{
		Name:        pgtype.Text{String: newName, Valid: true},
		Description: pgtype.Text{String: newDescription, Valid: true},
		CategoryID:  originalCategory.CategoryID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedCategory)

	require.Equal(t, newName, updatedCategory.Name)
	require.Equal(t, newDescription, updatedCategory.Description)

	// Clean up
	err = testQuerier.DeleteCategory(context.Background(), originalCategory.CategoryID)
	require.NoError(t, err)
}
