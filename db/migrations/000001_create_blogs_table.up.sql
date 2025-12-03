CREATE TABLE IF NOT EXISTS blogs(
    blog_id serial PRIMARY KEY,
    name VARCHAR (50) UNIQUE NOT NULL,
    content TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);