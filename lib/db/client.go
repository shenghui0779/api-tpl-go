package db

import (
	"context"
	"fmt"
	"time"

	"api/ent"
	"api/lib/log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/shenghui0779/yiigo"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cli *ent.Client

// Init 初始化Ent实例(如有多个实例，在此方法中初始化)
func Init() error {
	cfg := &yiigo.DBConfig{
		Driver: viper.GetString("db.driver"),
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
		return err
	}

	driver := entsql.OpenDB(dialect.MySQL, db)

	var iDriver dialect.Driver = driver
	if viper.GetBool("debug") {
		iDriver = dialect.DebugWithContext(driver, func(ctx context.Context, v ...any) {
			log.Info(ctx, "SQL info", zap.String("SQL", fmt.Sprint(v...)))
		})
	}
	cli = ent.NewClient(ent.Driver(iDriver))

	return nil
}

// Client 返回Ent实例
func Client() *ent.Client {
	return cli
}
