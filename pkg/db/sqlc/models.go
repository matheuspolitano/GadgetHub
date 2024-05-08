// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	CategoryID  int32  `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ChatMessage struct {
	ChatMessageID   int32              `json:"chat_message_id"`
	ChatSessionID   int32              `json:"chat_session_id"`
	MessageReceived pgtype.Text        `json:"message_received"`
	MessageSent     pgtype.Text        `json:"message_sent"`
	ReceivedAt      pgtype.Timestamptz `json:"received_at"`
	SentAt          pgtype.Timestamptz `json:"sent_at"`
	Action          string             `json:"action"`
	MessageBeforeID pgtype.Int4        `json:"message_before_id"`
}

type ChatSession struct {
	ChatSessionID int32              `json:"chat_session_id"`
	LastMessageID pgtype.Int4        `json:"last_message_id"`
	ActionFlow    string             `json:"action_flow"`
	UserID        int32              `json:"user_id"`
	Payload       string             `json:"payload"`
	OpenedAt      time.Time          `json:"opened_at"`
	ClosedAt      pgtype.Timestamptz `json:"closed_at"`
}

type DiscountCoupon struct {
	CouponID  int32       `json:"coupon_id"`
	CreatedBy int32       `json:"created_by"`
	CreatedAt pgtype.Date `json:"created_at"`
	ExpiresAt pgtype.Date `json:"expires_at"`
}

type Order struct {
	OrderID   int32          `json:"order_id"`
	ProductID int32          `json:"product_id"`
	UserID    int32          `json:"user_id"`
	CouponID  pgtype.Int4    `json:"coupon_id"`
	Price     pgtype.Numeric `json:"price"`
	CreatedAt time.Time      `json:"created_at"`
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
	OrderID    int32       `json:"order_id"`
	CreatedAt  time.Time   `json:"created_at"`
	Rating     int32       `json:"rating"`
	ReviewDate pgtype.Date `json:"review_date"`
}

type User struct {
	UserID       int32     `json:"user_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	HashPassword string    `json:"hash_password"`
	Phone        string    `json:"phone"`
	UserRole     string    `json:"user_role"`
}
