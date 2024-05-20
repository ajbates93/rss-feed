CREATE TABLE IF NOT EXISTS feeds (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS feed_items (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    feed_id INTEGER NOT NULL,
    author VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    url VARCHAR(255) NOT NULL,
    published_at TIMESTAMP,
    image_url VARCHAR(255),
    FOREIGN KEY (feed_id) REFERENCES feeds (id)
);
