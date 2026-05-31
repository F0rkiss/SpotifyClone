CREATE TABLE
    users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        email VARCHAR(320) NOT NULL UNIQUE,
        username VARCHAR(320) NOT NULL UNIQUE,
        hashed_password VARCHAR(96) NOT NULL,
        birth_date DATE,
        age VARCHAR(50),
        gender VARCHAR(50),
        country VARCHAR(100),
        created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    refresh_token (
        user_id UUID REFERENCES users (id) ON DELETE CASCADE,
        hashed_token VARCHAR(500) NOT NULL UNIQUE,
        created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
        expires_at TIMESTAMPTZ NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '1 day'),
        PRIMARY KEY (user_id, hashed_token)
    );