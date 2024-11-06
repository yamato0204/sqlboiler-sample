CREATE TABLE IF NOT EXISTS recipes (
    id CHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    thumbnail_url VARCHAR(255),
    recipe TEXT NOT NULL,
    category_id INT NOT NULL, -- Referencing category table
    ingredient JSON NOT NULL, -- Keeping ingredients as JSON
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_recipe_category FOREIGN KEY (category_id) REFERENCES categories(id) -- Adding the foreign key constraint
);
