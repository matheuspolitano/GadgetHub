package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

// Helper function to generate review parameters
func generateReviewParams(orderID int32) CreateReviewParams {
	return CreateReviewParams{
		OrderID:    orderID,
		Rating:     5, // Assuming a rating scale of 1-5
		ReviewDate: pgtype.Date{Time: time.Now(), Valid: true},
	}
}

// Helper function to create a review for testing
func createReviewTest(t *testing.T, orderID int32) Review {
	reviewParams := generateReviewParams(orderID)
	review, err := testQuerier.CreateReview(context.TODO(), reviewParams)
	require.NoError(t, err)
	require.NotEmpty(t, review)
	return review
}
func TestCreateReview(t *testing.T) {
	// Set up necessary data
	user := createUserTest(t)
	category := createCategoryTest(t)
	product := createProductTest(t, category.CategoryID)
	order := createOrderTest(t, user.UserID, product.ProductID)

	// Create a review
	reviewParams := generateReviewParams(order.OrderID)
	review, err := testQuerier.CreateReview(context.TODO(), reviewParams)
	require.NoError(t, err)
	require.NotEmpty(t, review)

	// Validate the review data
	require.Equal(t, reviewParams.OrderID, review.OrderID)
	require.Equal(t, reviewParams.Rating, review.Rating)
	require.Equal(t, reviewParams.ReviewDate.Time.Day(), review.ReviewDate.Time.Day()) // Check date equality at day level

	// Clean up test data
	err = testQuerier.DeleteReview(context.Background(), review.ReviewID)
	require.NoError(t, err)
	testQuerier.DeleteOrder(context.Background(), order.OrderID)
	testQuerier.DeleteProduct(context.Background(), product.ProductID)
	testQuerier.DeleteCategory(context.Background(), category.CategoryID)
	testQuerier.DeleteUser(context.Background(), user.UserID)
}

func TestGetReviewsByOrder(t *testing.T) {
	// Set up necessary data
	user := createUserTest(t)
	category := createCategoryTest(t)
	product := createProductTest(t, category.CategoryID)
	order := createOrderTest(t, user.UserID, product.ProductID)
	review1 := createReviewTest(t, order.OrderID)
	review2 := createReviewTest(t, order.OrderID)

	// Fetch reviews by order ID
	reviews, err := testQuerier.GetReviewsByOrder(context.TODO(), order.OrderID)
	require.NoError(t, err)
	require.Len(t, reviews, 2)

	// Check that the reviews fetched match those created
	for _, review := range reviews {
		require.True(t, review.ReviewID == review1.ReviewID || review.ReviewID == review2.ReviewID)
	}

	// Clean up test data
	testQuerier.DeleteReview(context.Background(), review1.ReviewID)
	testQuerier.DeleteReview(context.Background(), review2.ReviewID)
	testQuerier.DeleteOrder(context.Background(), order.OrderID)
	testQuerier.DeleteProduct(context.Background(), product.ProductID)
	testQuerier.DeleteCategory(context.Background(), category.CategoryID)
	testQuerier.DeleteUser(context.Background(), user.UserID)
}
