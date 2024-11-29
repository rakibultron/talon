-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Unique identifier for the user
    name VARCHAR(255) NOT NULL,                   -- User's full name
    email VARCHAR(255) UNIQUE NOT NULL,           -- User's unique email address
    password TEXT NOT NULL,                  -- Hashed password for security
    role VARCHAR(50) NOT NULL DEFAULT 'customer', -- User role (e.g., admin, vendor, customer)
    is_active BOOLEAN NOT NULL DEFAULT TRUE,      -- Indicates if the account is active
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(), -- Timestamp for when the user was created
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()  -- Timestamp for when the user was last updated
);

CREATE UNIQUE INDEX idx_users_email ON users(email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
