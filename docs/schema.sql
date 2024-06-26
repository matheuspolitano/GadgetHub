-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2024-05-08T08:27:22.262Z

CREATE TABLE "Users" (
  "user_id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "hash_password" varchar NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "user_role" varchar NOT NULL
);

CREATE TABLE "DiscountCoupons" (
  "coupon_id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_by" int NOT NULL,
  "created_at" date NOT NULL,
  "expires_at" date NOT NULL
);

CREATE TABLE "Categories" (
  "category_id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL
);

CREATE TABLE "Products" (
  "product_id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" decimal NOT NULL,
  "stock" int NOT NULL,
  "category_id" int NOT NULL,
  "brand" varchar,
  "model" varchar
);

CREATE TABLE "Order" (
  "order_id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "product_id" int NOT NULL,
  "user_id" int NOT NULL,
  "coupon_id" int,
  "price" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Reviews" (
  "review_id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "order_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "rating" int NOT NULL,
  "review_date" date NOT NULL
);

CREATE TABLE "ChatMessage" (
  "chat_message_id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "chat_session_id" int NOT NULL,
  "message_received" varchar,
  "message_sent" varchar,
  "received_at" timestamptz,
  "sent_at" timestamptz,
  "action" varchar NOT NULL,
  "message_before_id" int
);

CREATE TABLE "ChatSession" (
  "chat_session_id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "last_message_id" int,
  "action_flow" varchar NOT NULL,
  "user_id" int NOT NULL,
  "payload" varchar NOT NULL,
  "opened_at" timestamptz NOT NULL DEFAULT (now()),
  "closed_at" timestamptz
);

CREATE INDEX ON "Products" ("category_id");

CREATE INDEX ON "Order" ("order_id");

CREATE INDEX ON "Order" ("product_id");

CREATE INDEX ON "Reviews" ("order_id");

CREATE INDEX ON "ChatMessage" ("chat_session_id");

CREATE INDEX ON "ChatSession" ("user_id");

ALTER TABLE "DiscountCoupons" ADD FOREIGN KEY ("created_by") REFERENCES "Users" ("user_id");

ALTER TABLE "Products" ADD FOREIGN KEY ("category_id") REFERENCES "Categories" ("category_id");

ALTER TABLE "Order" ADD FOREIGN KEY ("product_id") REFERENCES "Products" ("product_id");

ALTER TABLE "Order" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("user_id");

ALTER TABLE "Order" ADD FOREIGN KEY ("coupon_id") REFERENCES "DiscountCoupons" ("coupon_id");

ALTER TABLE "Reviews" ADD FOREIGN KEY ("order_id") REFERENCES "Order" ("order_id");

ALTER TABLE "ChatMessage" ADD FOREIGN KEY ("chat_session_id") REFERENCES "ChatSession" ("chat_session_id");

ALTER TABLE "ChatMessage" ADD FOREIGN KEY ("message_before_id") REFERENCES "ChatMessage" ("chat_message_id");

ALTER TABLE "ChatSession" ADD FOREIGN KEY ("last_message_id") REFERENCES "ChatMessage" ("chat_message_id");

ALTER TABLE "ChatSession" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("user_id");
