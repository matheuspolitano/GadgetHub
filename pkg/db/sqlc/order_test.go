package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// Assuming you have helper functions like createUserTest and createProductTest already implemented

func generateOrderParams(userID, productID int32) CreateOrderParams {
	return CreateOrderParams{
		ProductID: productID,
		UserID:    userID,
		Price:     99.99,
	}
}

func createOrderTest(t *testing.T, userID, productID int32) Order {
	orderParams := generateOrderParams(userID, productID)
	order, err := testQuerier.CreateOrder(context.TODO(), orderParams)
	require.NoError(t, err)
	require.NotEmpty(t, order)
	return order
}
func TestCreateOrder(t *testing.T) {
	user := createUserTest(t)                            // Create a user for the order
	category := createCategoryTest(t)                    // Create a category for the product
	product := createProductTest(t, category.CategoryID) // Create a product for the order

	orderParams := generateOrderParams(user.UserID, product.ProductID)
	order, err := testQuerier.CreateOrder(context.TODO(), orderParams)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, orderParams.UserID, order.UserID)
	require.Equal(t, orderParams.ProductID, order.ProductID)
	require.Equal(t, orderParams.Price, order.Price)

	// Clean up
	err = testQuerier.DeleteOrder(context.Background(), order.OrderID)
	require.NoError(t, err)
	testQuerier.DeleteProduct(context.Background(), product.ProductID)
	testQuerier.DeleteCategory(context.Background(), category.CategoryID)
	testQuerier.DeleteUser(context.Background(), user.UserID)
}
func TestGetOrder(t *testing.T) {
	user := createUserTest(t)
	category := createCategoryTest(t)
	product := createProductTest(t, category.CategoryID)
	order := createOrderTest(t, user.UserID, product.ProductID)

	fetchedOrder, err := testQuerier.GetOrder(context.TODO(), order.OrderID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedOrder)

	require.Equal(t, order.OrderID, fetchedOrder.OrderID)

	// Clean up
	err = testQuerier.DeleteOrder(context.Background(), order.OrderID)
	require.NoError(t, err)
	testQuerier.DeleteProduct(context.Background(), product.ProductID)
	testQuerier.DeleteCategory(context.Background(), category.CategoryID)
	testQuerier.DeleteUser(context.Background(), user.UserID)
}
func TestGetOrdersByProduct(t *testing.T) {
	user := createUserTest(t)
	category := createCategoryTest(t)
	product := createProductTest(t, category.CategoryID)
	order1 := createOrderTest(t, user.UserID, product.ProductID)
	order2 := createOrderTest(t, user.UserID, product.ProductID)

	orders, err := testQuerier.GetOrdersByProduct(context.TODO(), product.ProductID)
	require.NoError(t, err)
	require.NotEmpty(t, orders)
	require.True(t, len(orders) >= 2) // Check for at least two orders

	// Clean up
	testQuerier.DeleteOrder(context.Background(), order1.OrderID)
	testQuerier.DeleteOrder(context.Background(), order2.OrderID)
	testQuerier.DeleteProduct(context.Background(), product.ProductID)
	testQuerier.DeleteCategory(context.Background(), category.CategoryID)
	testQuerier.DeleteUser(context.Background(), user.UserID)
}

func TestGetOrdersByUser(t *testing.T) {
	user := createUserTest(t)                             // Create a user
	category := createCategoryTest(t)                     // Create a category for the product
	product1 := createProductTest(t, category.CategoryID) // Create first product for the order
	product2 := createProductTest(t, category.CategoryID) // Create second product for the order

	// Create two orders for the same user but different products
	order1 := createOrderTest(t, user.UserID, product1.ProductID)
	order2 := createOrderTest(t, user.UserID, product2.ProductID)

	// Fetch orders by user ID
	orders, err := testQuerier.GetOrdersByUser(context.TODO(), user.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, orders)
	require.True(t, len(orders) >= 2) // Check for at least two orders for the user

	// Assertions to confirm that fetched orders belong to the user
	for _, order := range orders {
		require.Equal(t, user.UserID, order.UserID)
	}

	// Clean up
	testQuerier.DeleteOrder(context.Background(), order1.OrderID)
	testQuerier.DeleteOrder(context.Background(), order2.OrderID)
	testQuerier.DeleteProduct(context.Background(), product1.ProductID)
	testQuerier.DeleteProduct(context.Background(), product2.ProductID)
	testQuerier.DeleteCategory(context.Background(), category.CategoryID)
	testQuerier.DeleteUser(context.Background(), user.UserID)
}
