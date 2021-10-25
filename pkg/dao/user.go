package dao

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo"

	"tplgo/pkg/logger"
	"tplgo/pkg/models"
)

type UserDao interface {
	FindByID(ctx context.Context, id int64) (*models.User, error)
}

func NewUserDao() UserDao {
	return &user{
		db:      yiigo.DB(),
		table:   "t_user",
		builder: yiigo.NewMySQLBuilder(yiigo.WithSQLLogger(logger.SQL)),
	}
}

type user struct {
	db      *sqlx.DB
	table   string
	builder yiigo.SQLBuilder
}

func (u *user) FindByID(ctx context.Context, id int64) (*models.User, error) {
	query, binds := u.builder.Wrap(
		yiigo.Table(u.table),
		yiigo.Select("id", "nickname", "avatar", "phone"),
		yiigo.Where("id = ?", id),
		yiigo.Limit(1),
	).ToQuery(ctx)

	record := new(models.User)

	if err := u.db.Get(record, query, binds...); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "Dao.User.FindByID error")
		}

		return nil, nil
	}

	return record, nil
}
