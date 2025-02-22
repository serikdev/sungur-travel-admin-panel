-- +goose Up
CREATE TABLE hotels(
    ID SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    stars INTEGER NOT NULL CHECK(stars BETWEEN 1 AND 5),
    address TEXT NOT NULL,
    images TEXT[],
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE hotels;

