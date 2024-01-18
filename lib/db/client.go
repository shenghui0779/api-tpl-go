package db

import (
	"api/ent"
	"api/lib/log"
	"time"

	"context"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/shenghui0779/yiigo/db"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cli *ent.Client

// Init 初始化Ent实例(如有多个实例，在此方法中初始化)
func Init() error {
	cfg := buildCfg(viper.GetString("db.driver"), viper.GetString("db.dsn"), viper.GetStringMap("db.options"))

	db, err := db.New(cfg)
	if err != nil {
		return err
	}

	cli = ent.NewClient(
		ent.Driver(dialect.DebugWithContext(
			entsql.OpenDB(cfg.Driver, db),
			func(ctx context.Context, v ...any) {
				log.Info(ctx, "SQL info", zap.String("SQL", fmt.Sprint(v...)))
			}),
		),
	)

	return nil
}

// Client 返回Ent实例
func Client() *ent.Client {
	return cli
}

func buildCfg(driver, dsn string, opts map[string]any) *db.Config {
	cfg := &db.Config{
		Driver: driver,
		DSN:    dsn,
	}

	if len(opts) != 0 {
		cfg.Options = &db.Options{
			MaxOpenConns:    cast.ToInt(opts["max_open_conns"]),
			MaxIdleConns:    cast.ToInt(opts["max_idle_conns"]),
			ConnMaxLifetime: cast.ToDuration(opts["conn_max_lifetime"]) * time.Second,
			ConnMaxIdleTime: cast.ToDuration(opts["conn_max_idle_time"]) * time.Second,
		}
	}

	return cfg
}
