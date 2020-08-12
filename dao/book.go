package dao

import (
	"database/sql"

	"github.com/shenghui0779/yiigo"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/shenghui0779/demo/models"
)

type Book struct {
	db  *sqlx.DB
	orm *gorm.DB
}

func NewBook() *Book {
	return &Book{
		db:  yiigo.DB(),
		orm: yiigo.Orm(),
	}
}

func (b *Book) FindByID(id int64) (*models.Book, error) {
	record := new(models.Book)

	if err := b.db.Get(record, "SELECT * FROM `book` WHERE `id` = ? LIMIT 1", id); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "find book by id error")
		}

		return nil, nil
	}

	return record, nil
}
