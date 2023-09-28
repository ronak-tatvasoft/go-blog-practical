CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    nickname VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content text NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);