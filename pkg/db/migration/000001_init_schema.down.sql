-- Drop Foreign Key Constraints
ALTER TABLE Order_Items DROP CONSTRAINT IF EXISTS order_items_order_id_fkey;
ALTER TABLE Order_Items DROP CONSTRAINT IF EXISTS order_items_product_id_fkey;
ALTER TABLE Order_Items DROP CONSTRAINT IF EXISTS order_items_coupon_id_fkey;

ALTER TABLE Orders DROP CONSTRAINT IF EXISTS orders_user_id_fkey;

ALTER TABLE Cart_Items DROP CONSTRAINT IF EXISTS cart_items_user_id_fkey;
ALTER TABLE Cart_Items DROP CONSTRAINT IF EXISTS cart_items_product_id_fkey;

ALTER TABLE Reviews DROP CONSTRAINT IF EXISTS reviews_product_id_fkey;
ALTER TABLE Reviews DROP CONSTRAINT IF EXISTS reviews_user_id_fkey;

ALTER TABLE Products DROP CONSTRAINT IF EXISTS products_category_id_fkey;

ALTER TABLE DiscountCoupons DROP CONSTRAINT IF EXISTS discount_coupons_created_by_fkey;

-- Drop Tables
DROP TABLE IF EXISTS Order_Items;
DROP TABLE IF EXISTS Orders;
DROP TABLE IF EXISTS Cart_Items;
DROP TABLE IF EXISTS Reviews;
DROP TABLE IF EXISTS Products;
DROP TABLE IF EXISTS Categories;
DROP TABLE IF EXISTS DiscountCoupons;
DROP TABLE IF EXISTS Users;
