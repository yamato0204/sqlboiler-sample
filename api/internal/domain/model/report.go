package model

type Report struct {
	ID           string   `db:"id"`
	Comment      string   `db:"comment"`
	ThumbnailUrl string   `db:"thumbnail_url"`
	UserID       string   `db:"user_id"`
	Recipe       Recipe   `db:"recipe"`
	Category     Category `db:"category"`
	CreatedAt    string   `db:"created_at"`
	UpdatedAt    string   `db:"updated_at"`
}
