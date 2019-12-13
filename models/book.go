package models

// Book model for query
type Book struct {
	ID          int64   `db:"id" gorm:"column:id"`
	Title       string  `db:"title" gorm:"column:title"`
	SubTitle    string  `db:"subtitle" gorm:"column:subtitle"`
	Author      string  `db:"author" gorm:"column:author"`
	Version     string  `db:"version" gorm:"column:version"`
	Price       float64 `db:"price" gorm:"column:price"`
	Publisher   string  `db:"publisher" gorm:"column:publisher"`
	PublishDate string  `db:"publish_date" gorm:"column:publish_date"`
	CreatedAt   int64   `db:"created_at" gorm:"column:created_at"`
	UpdatedAt   int64   `db:"updated_at" gorm:"column:updated_at"`
}

func (b *Book) TableName() string {
	return "book"
}
