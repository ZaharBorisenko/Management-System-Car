CREATE TABLE brands (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

ALTER TABLE cars
ADD COLUMN brand_id UUID;

ALTER TABLE engines
ADD COLUMN brand_id UUID;

