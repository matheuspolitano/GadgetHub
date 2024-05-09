// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateChatMessage(ctx context.Context, arg CreateChatMessageParams) (ChatMessage, error)
	CreateChatSession(ctx context.Context, arg CreateChatSessionParams) (ChatSession, error)
	CreateDiscountCoupon(ctx context.Context, arg CreateDiscountCouponParams) (DiscountCoupon, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCategory(ctx context.Context, categoryID int32) error
	DeleteChatMessage(ctx context.Context, chatMessageID int32) error
	DeleteChatSession(ctx context.Context, chatSessionID int32) error
	DeleteDiscountCoupon(ctx context.Context, couponID int32) error
	DeleteOrder(ctx context.Context, orderID int32) error
	DeleteProduct(ctx context.Context, productID int32) error
	DeleteReview(ctx context.Context, reviewID int32) error
	DeleteUser(ctx context.Context, userID int32) error
	GetCategoriesByName(ctx context.Context, name string) ([]Category, error)
	GetCategory(ctx context.Context, categoryID int32) (Category, error)
	GetChatMessage(ctx context.Context, chatMessageID int32) (ChatMessage, error)
	GetChatMessagesBySession(ctx context.Context, chatSessionID int32) ([]ChatMessage, error)
	GetChatSession(ctx context.Context, chatSessionID int32) (ChatSession, error)
	GetChatSessionsByUser(ctx context.Context, userID int32) ([]ChatSession, error)
	GetCouponsByCreator(ctx context.Context, createdBy int32) ([]DiscountCoupon, error)
	GetDiscountCoupon(ctx context.Context, couponID int32) (DiscountCoupon, error)
	GetOrder(ctx context.Context, orderID int32) (Order, error)
	GetOrdersByProduct(ctx context.Context, productID int32) ([]Order, error)
	GetOrdersByUser(ctx context.Context, userID int32) ([]Order, error)
	GetProduct(ctx context.Context, productID int32) (Product, error)
	GetProductsByCategory(ctx context.Context, categoryID int32) ([]Product, error)
	GetReview(ctx context.Context, reviewID int32) (Review, error)
	GetReviewsByOrder(ctx context.Context, orderID int32) ([]Review, error)
	GetUser(ctx context.Context, userID int32) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByPhone(ctx context.Context, phone string) (User, error)
	// Allows searching by name with case-insensitive matching.
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateChatMessage(ctx context.Context, arg UpdateChatMessageParams) (ChatMessage, error)
	UpdateChatSession(ctx context.Context, arg UpdateChatSessionParams) (ChatSession, error)
	UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error)
	UpdateReview(ctx context.Context, arg UpdateReviewParams) (Review, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
