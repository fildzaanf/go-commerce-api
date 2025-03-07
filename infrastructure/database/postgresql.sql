CREATE DATABASE gocommercedb;

\c gocommercedb;

CREATE TYPE user_role_enum AS ENUM ('user', 'buyer', 'seller');

CREATE TYPE payment_status_enum AS ENUM ('deny', 'success', 'cancel', 'expire', 'pending');

CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role user_role_enum DEFAULT 'user',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE products (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(15,2) DEFAULT 0.00 NOT NULL,
    stock INT NOT NULL,
    image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE payments (
    id VARCHAR(36) PRIMARY KEY,
    product_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36)NOT NULL,
    payment_code VARCHAR(36) NOT NULL UNIQUE,
    quantity INT NOT NULL,
    total_amount DECIMAL(15,2) NOT NULL,
    status payment_status_enum DEFAULT 'pending',
    payment_url TEXT,
    token TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);


CREATE INDEX idx_products_user_id ON products(user_id);
CREATE INDEX idx_payments_product_id ON payments(product_id);
