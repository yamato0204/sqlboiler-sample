CREATE TABLE IF NOT EXISTS reports (
    id VARCHAR(36) PRIMARY KEY,
    comment TEXT,
    thumbnail_url VARCHAR(255),
    user_id VARCHAR(36) NOT NULL,
    recipe_id CHAR(36) NOT NULL, -- レシピIDを追加
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_taberepo_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_taberepo_recipe FOREIGN KEY (recipe_id) REFERENCES recipes(id) -- レシピIDを外部キーとして参照
);
