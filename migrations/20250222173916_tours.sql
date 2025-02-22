-- +goose Up
CREATE TABLE tours(
    ID SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL CHECK(price > 0),
    duration INTEGER NOT NULL CHECK(duration > 0), 
    hotel_id INTEGER REFERENCES hotels(id) ON DELETE SET NULL ON UPDATE CASCADE,
    images TEXT[],
    start_dates DATE[],
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- +goose Down
DROP TABLE tours;