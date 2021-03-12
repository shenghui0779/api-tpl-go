package models

// Book model for query
type Book struct {
	ID          int64   `db:"id"`
	Title       string  `db:"title"`
	SubTitle    string  `db:"subtitle"`
	Author      string  `db:"author"`
	Version     string  `db:"version"`
	Price       float64 `db:"price"`
	Publisher   string  `db:"publisher"`
	PublishDate string  `db:"publish_date"`
	CreatedAt   int64   `db:"created_at"`
	UpdatedAt   int64   `db:"updated_at"`
}
