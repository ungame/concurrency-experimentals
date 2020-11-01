DROP TABLE IF EXISTS products;

CREATE TABLE products (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    description TEXT,
    quantity INT DEFAULT 0,
    price NUMERIC(18, 2),
    available CHAR(1) DEFAULT '0',
    photo_url  VARCHAR(255),
    ratings INT DEFAULT 0,
    category VARCHAR(255),
    manufacturer VARCHAR(255),
    is_deleted CHAR(1) DEFAULT '0',
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP NULL
);

\d+ products;
