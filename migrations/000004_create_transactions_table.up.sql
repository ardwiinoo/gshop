CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_public_id varchar(100) NOT NULL,
    product_id int NOT NULL,
    product_price int NOT NULL,
    amount int NOT NULL,
    sub_total int NOT NULL,
    platform_fee int DEFAULT 0,
    grand_total int NOT NULL,
    status int NOT NULL,
    product_snapshot jsonb,
    created_at timestamp,
    updated_at timestamp
)