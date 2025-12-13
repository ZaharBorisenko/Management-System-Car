ALTER TABLE cars
    ADD CONSTRAINT fk_cars_brand
        FOREIGN KEY (brand_id) REFERENCES brands(id);

ALTER TABLE engines
    ADD CONSTRAINT fk_engines_brand
        FOREIGN KEY (brand_id) REFERENCES brands(id);

CREATE INDEX idx_cars_brand_id ON cars(brand_id);
CREATE INDEX idx_engines_brand_id ON engines(brand_id);