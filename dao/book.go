package dao

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/shenghui0779/demo/models"
	"github.com/shenghui0779/yiigo"
)

type BookDao interface {
	FindByID(id int64) (*models.Book, error)
}

type book struct {
	db      *sqlx.DB
	table   string
	builder yiigo.SQLBuilder
}

func NewBookDao() BookDao {
	return &book{
		db:      yiigo.DB(),
		table:   "book",
		builder: yiigo.NewSQLBuilder(yiigo.MySQL),
	}
}

func (b *book) FindByID(id int64) (*models.Book, error) {
	query, binds := b.builder.Wrap(
		yiigo.Table(b.table),
		yiigo.Where("id = ?", id),
		yiigo.Limit(1),
	).ToQuery()

	record := new(models.Book)

	if err := b.db.Get(record, query, binds...); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "Dao.Book.FindByID error")
		}

		return nil, nil
	}

	return record, nil
}
