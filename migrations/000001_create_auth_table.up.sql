CREATE TABLE auth (
    id SERIAL PRIMARY KEY,
    email varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    role varchar(20) NOT NULL DEFAULT 'user',
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now()
)