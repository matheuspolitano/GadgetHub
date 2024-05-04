// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CartItem struct {
	CartItemID int32 `json:"cart_item_id"`
	UserID     int32 `json:"user_id"`
	ProductID  int32 `json:"product_id"`
	Quantity   int32 `json:"quantity"`
}

type Category struct {
	CategoryID  int32  `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DiscountCoupon struct {
	CouponID   int32       `json:"coupon_id"`
	CreatedBy  int32       `json:"created_by"`
	CreatedAt  pgtype.Date `json:"created_at"`
	ValidUntil pgtype.Date `json:"valid_until"`
}

type Order struct {
	OrderID     int32          `json:"order_id"`
	UserID      int32          `json:"user_id"`
	OrderDate   pgtype.Date    `json:"order_date"`
	Status      string         `json:"status"`
	TotalAmount pgtype.Numeric `json:"total_amount"`
}

type OrderItem struct {
	OrderItemID int32          `json:"order_item_id"`
	OrderID     int32          `json:"order_id"`
	ProductID   int32          `json:"product_id"`
	CouponID    pgtype.Int4    `json:"coupon_id"`
	Quantity    int32          `json:"quantity"`
	UnitPrice   pgtype.Numeric `json:"unit_price"`
}

type Product struct {
	ProductID   int32          `json:"product_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       pgtype.Numeric `json:"price"`
	Stock       int32          `json:"stock"`
	CategoryID  int32          `json:"category_id"`
	Brand       pgtype.Text    `json:"brand"`
	Model       pgtype.Text    `json:"model"`
}

type Review struct {
	ReviewID   int32       `json:"review_id"`
	ProductID  int32       `json:"product_id"`
	UserID     int32       `json:"user_id"`
	Rating     int32       `json:"rating"`
	ReviewDate pgtype.Date `json:"review_date"`
}

type User struct {
	UserID       int32  `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
	Phone        string `json:"phone"`
	IsAdmin      bool   `json:"is_admin"`
}
