package models

import "time"

// Book model for query
type Book struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	SubTitle    string    `db:"subtitle" json:"subtitle"`
	Author      string    `db:"author" json:"author"`
	Version     string    `db:"version" json:"version"`
	Price       string    `db:"price" json:"price"`
	Publisher   string    `db:"publisher" json:"publisher"`
	PublishDate string    `db:"publish_date" json:"publish_date"`
	CreatedAt   time.Time `db:"created_at" json:"-"`
	UpdatedAt   time.Time `db:"updated_at" json:"-"`
}

// BookAdd model for add
type BookAdd struct {
	Title       string    `db:"title" form:"title" binding:"required"`
	SubTitle    string    `db:"subtitle" form:"subtitle"`
	Author      string    `db:"author" form:"author" binding:"required"`
	Version     string    `db:"version" form:"version" binding:"required"`
	Price       string    `db:"price" form:"price" binding:"required"`
	Publisher   string    `db:"publisher" form:"publisher" binding:"required"`
	PublishDate string    `db:"publish_date" form:"publish_date" binding:"required"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// BookEdit model for edit
type BookEdit struct {
	Title       string    `db:"title" form:"title" binding:"required"`
	SubTitle    string    `db:"subtitle" form:"subtitle"`
	Author      string    `db:"author" form:"author" binding:"required"`
	Version     string    `db:"version" form:"version" binding:"required"`
	Price       string    `db:"price" form:"price" binding:"required"`
	Publisher   string    `db:"publisher" form:"publisher" binding:"required"`
	PublishDate string    `db:"publish_date" form:"publish_date" binding:"required"`
	UpdatedAt   time.Time `db:"updated_at"`
}
