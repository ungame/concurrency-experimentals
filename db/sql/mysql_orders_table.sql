DROP TABLE IF EXISTS orders;

CREATE TABLE orders (
    id VARCHAR(100) PRIMARY KEY,
    user_id VARCHAR(100),
    product_id VARCHAR(100),
    description TEXT,
    quantity INT,
    unit_price DECIMAL(18, 2),
    amount DECIMAL(18, 2),
    status VARCHAR(35),
    reason_reject VARCHAR(255),
    paid CHAR(1) DEFAULT '0',
    payment_type VARCHAR(35),
    is_deleted CHAR(1) DEFAULT '0',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    deleted_at TIMESTAMP NULL,
    canceled_at TIMESTAMP  NULL
);

desc orders;