-- ===========================
-- ENGINE TABLE
-- ===========================
CREATE TABLE engine (
    id UUID PRIMARY KEY,
    description TEXT NOT NULL,
    displacement BIGINT NOT NULL CHECK (displacement > 0),
    no_of_cylinders BIGINT NOT NULL CHECK (no_of_cylinders >= 1),
    car_range BIGINT NOT NULL CHECK (car_range >= 0),
    horse_power BIGINT NOT NULL CHECK (horse_power >= 0),
    torque BIGINT NOT NULL CHECK (torque >= 0),
    engine_type TEXT NOT NULL CHECK (engine_type IN ('Turbine', 'electric', 'rotary', 'Hybrid', 'petrol', 'diesel')),
    emission_class TEXT NOT NULL CHECK (  emission_class IN ('Euro 3', 'Euro 4', 'Euro 5', 'Euro 6', 'Euro 6d')),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


-- ===========================
-- CAR TABLE
-- ===========================
CREATE TABLE car (
    id UUID PRIMARY KEY,
    description TEXT NOT NULL,
    year BIGINT NOT NULL CHECK (year >= 1886 AND year <= 2100),
    brand TEXT NOT NULL,
    model TEXT NOT NULL,
    fuel_type TEXT NOT NULL CHECK (fuel_type IN ('petrol', 'diesel', 'electric', 'hybrid', 'hydrogen', 'gas')),
    engine_id UUID NOT NULL REFERENCES engine(id) ON DELETE RESTRICT,
    price NUMERIC(12,2) NOT NULL CHECK (price >= 0),
    vin CHAR(17) NOT NULL UNIQUE,
    mileage BIGINT NOT NULL CHECK (mileage >= 0),
    transmission TEXT NOT NULL CHECK (transmission IN ('automatic', 'manual', 'robotic', 'cvt')),
    color TEXT NOT NULL,
    body_type TEXT NOT NULL CHECK (body_type IN ('sedan', 'hatchback', 'coupe', 'cabriolet','suv', 'crossover', 'pickup', 'wagon', 'van', 'micro')),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Index
CREATE INDEX idx_car_brand ON car(brand);
CREATE INDEX idx_car_model ON car(model);
CREATE INDEX idx_car_fuel ON car(fuel_type);
CREATE INDEX idx_car_engine_id ON car(engine_id);
