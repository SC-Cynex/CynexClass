CREATE TABLE IF NOT EXISTS TBLAddress (
    id SERIAL PRIMARY KEY,
    street VARCHAR(100) NOT NULL,          -- Street names typically do not exceed 100 characters
    number INTEGER NOT NULL,               -- House/building numbers are stored as integers
    neighborhood VARCHAR(100) NOT NULL,    -- Neighborhood names are usually short, within 100 characters
    city VARCHAR(100) NOT NULL,            -- City names typically do not exceed 100 characters
    state VARCHAR(2) NOT NULL,             -- State codes use two characters
    zip_code VARCHAR(8) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
