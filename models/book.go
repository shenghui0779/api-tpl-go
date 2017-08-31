package models

import "time"

type Book struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	SubTitle    string    `db:"subtitle" json:"subtitle"`
	Author      string    `db:"author" json:"author"`
	Version     string    `db:"version" json:"version"`
	Price       string    `db:"price" json:"price"`
	Publisher   string    `db:"publisher" json:"publisher"`
	PublishDate string    `db:"publish_date" json:"publish_date"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
