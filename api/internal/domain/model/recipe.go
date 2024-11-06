package model

type Recipe struct {
	ID           string   `db:"id"`
	Title        string   `db:"title"`
	ThumbnailUrl string   `db:"thumbnail_url"`
	Recipe       string   `db:"recipe"`
	Category     Category `db:"category"`
	Ingredient   string   `db:"ingredient"`
	CreatedAt    string   `db:"created_at"`
	UpdatedAt    string   `db:"updated_at"`
}
