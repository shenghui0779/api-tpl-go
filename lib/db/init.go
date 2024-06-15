package db

import (
	"context"
	"fmt"
	"time"

	"api/lib/log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/shenghui0779/yiigo"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Init 初始化Ent实例(如有多个实例，在此方法中初始化)
func Init() (dialect.Driver, error) {
	cfg := &yiigo.DBConfig{
		Driver: "pgx",
		DSN:    viper.GetString("db.dsn"),
	}

	opts := viper.GetStringMap("db.options")
	if len(opts) != 0 {
		cfg.MaxOpenConns = cast.ToInt(opts["max_open_conns"])
		cfg.MaxIdleConns = cast.ToInt(opts["max_idle_conns"])
		cfg.ConnMaxLifetime = cast.ToDuration(opts["conn_max_lifetime"]) * time.Second
		cfg.ConnMaxIdleTime = cast.ToDuration(opts["conn_max_idle_time"]) * time.Second
	}

	db, err := yiigo.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	driver := entsql.OpenDB(dialect.Postgres, db)
	if viper.GetBool("app.debug") {
		return dialect.DebugWithContext(driver, func(ctx context.Context, v ...any) {
			log.Info(ctx, "SQL info", zap.String("SQL", fmt.Sprint(v...)))
		}), nil
	}
	return driver, nil
}
