package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/matheuspolitano/GadgetHub/utils"
	"github.com/stretchr/testify/require"
)

func generateProductParams(categoryID int32) CreateProductParams {
	return CreateProductParams{
		Name:        utils.RandString(10),
		Description: utils.RandString(50),
		Price:       1.2,
		Stock:       100,
		CategoryID:  categoryID,
		Brand:       pgtype.Text{String: "TestBrand", Valid: true},
		Model:       pgtype.Text{String: "TestModel", Valid: true},
	}
}

func CreateProductTest(t *testing.T, categoryID int32) Product {
	prodParams := generateProductParams(categoryID)
	product, err := testQuerier.CreateProduct(context.TODO(), prodParams)
	require.NoError(t, err)
	require.NotEmpty(t, product)
	return product
}
func TestCreateProduct(t *testing.T) {
	category := createCategoryTest(t) // Assuming you have a similar helper for categories
	prodParams := generateProductParams(category.CategoryID)
	product, err := testQuerier.CreateProduct(context.TODO(), prodParams)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, prodParams.Name, product.Name)
	require.Equal(t, prodParams.Description, product.Description)
	require.Equal(t, prodParams.Price, product.Price)
	require.Equal(t, prodParams.Stock, product.Stock)
	require.Equal(t, prodParams.CategoryID, product.CategoryID)
	require.Equal(t, prodParams.Brand.String, product.Brand.String)
	require.Equal(t, prodParams.Model.String, product.Model.String)

	// Clean up
	err = testQuerier.DeleteProduct(context.Background(), product.ProductID)
	require.NoError(t, err)
	testQuerier.DeleteCategory(context.Background(), category.CategoryID)
}
func TestGetProduct(t *testing.T) {
	category := createCategoryTest(t)
	product := CreateProductTest(t, category.CategoryID)

	fetchedProduct, err := testQuerier.GetProduct(context.TODO(), product.ProductID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedProduct)

	// Assertions can be expanded based on fields
	require.Equal(t, product.Name, fetchedProduct.Name)

	// Clean up
	err = testQuerier.DeleteProduct(context.Background(), product.ProductID)
	require.NoError(t, err)
	testQuerier.DeleteCategory(context.Background(), category.CategoryID)
}
func TestGetProductsByCategory(t *testing.T) {
	category := createCategoryTest(t)
	product1 := CreateProductTest(t, category.CategoryID)
	product2 := CreateProductTest(t, category.CategoryID)

	products, err := testQuerier.GetProductsByCategory(context.TODO(), category.CategoryID)
	require.NoError(t, err)
	require.NotEmpty(t, products)
	require.True(t, len(products) >= 2) // Check for at least two products

	// Clean up
	testQuerier.DeleteProduct(context.Background(), product1.ProductID)
	testQuerier.DeleteProduct(context.Background(), product2.ProductID)
	testQuerier.DeleteCategory(context.Background(), category.CategoryID)
}
func TestUpdateProduct(t *testing.T) {
	category := createCategoryTest(t)
	product := CreateProductTest(t, category.CategoryID)
	newDescription := utils.RandString(50)

	updatedProduct, err := testQuerier.UpdateProduct(context.TODO(), UpdateProductParams{
		ProductID:   product.ProductID,
		Description: pgtype.Text{String: newDescription, Valid: true},
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedProduct)
	require.Equal(t, newDescription, updatedProduct.Description)

	// Clean up
	err = testQuerier.DeleteProduct(context.Background(), product.ProductID)
	require.NoError(t, err)
	testQuerier.DeleteCategory(context.Background(), category.CategoryID)
}
