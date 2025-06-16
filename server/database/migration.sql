CREATE TABLE establishments (
    id SERIAL PRIMARY KEY,
    number VARCHAR(20) NOT NULL,
    name VARCHAR(100) NOT NULL,
    corporate_name VARCHAR(100),
    address VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(2),
    zip_code VARCHAR(10),
    address_number VARCHAR(10)
);

CREATE TABLE stores (
    id SERIAL PRIMARY KEY,
    number VARCHAR(20) NOT NULL,
    name VARCHAR(100) NOT NULL,
    corporate_name VARCHAR(100),
    address VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(2),
    zip_code VARCHAR(10),
    address_number VARCHAR(10),
    establishment_id INT NOT NULL REFERENCES establishments(id) ON DELETE RESTRICT
);
