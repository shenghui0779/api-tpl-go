package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"api/ent"
	"api/lib/config"
	"api/lib/log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var cli *ent.Client

// Config 数据库初始化配置
type Config struct {
	// Driver 驱动名称
	Driver string

	// DSN 数据源名称
	// [-- MySQL] username:password@tcp(localhost:3306)/dbname?timeout=10s&charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local
	// [Postgres] host=localhost port=5432 user=root password=secret dbname=test connect_timeout=10 sslmode=disable
	// [- SQLite] file::memory:?cache=shared
	DSN string `json:"dsn"`

	// Options 配置选项
	Options *Options `json:"options"`
}

// Options 数据库配置选项
type Options struct {
	// MaxOpenConns 设置最大可打开的连接数；-1：不限；默认：20
	MaxOpenConns int `json:"max_open_conns"`

	// MaxIdleConns 连接池最大闲置连接数；-1：不限；默认：10
	MaxIdleConns int `json:"max_idle_conns"`

	// ConnMaxLifetime 连接的最大生命时长；-1：不限；默认：10分钟
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`

	// ConnMaxIdleTime 连接最大闲置时间；-1：不限；默认：5分钟
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time"`
}

func (o *Options) rebuild(opt *Options) {
	if opt.MaxOpenConns > 0 {
		o.MaxOpenConns = opt.MaxOpenConns
	} else {
		if opt.MaxOpenConns == -1 {
			o.MaxOpenConns = 0
		}
	}

	if opt.MaxIdleConns > 0 {
		o.MaxIdleConns = opt.MaxIdleConns
	} else {
		if opt.MaxIdleConns == -1 {
			o.MaxIdleConns = 0
		}
	}

	if opt.ConnMaxLifetime > 0 {
		o.ConnMaxLifetime = opt.ConnMaxLifetime
	} else {
		if opt.ConnMaxLifetime == -1 {
			o.ConnMaxLifetime = 0
		}
	}

	if opt.ConnMaxIdleTime > 0 {
		o.ConnMaxIdleTime = opt.ConnMaxIdleTime
	} else {
		if opt.ConnMaxIdleTime == -1 {
			o.ConnMaxIdleTime = 0
		}
	}
}

// Init 初始化Ent实例
func Init(cfg *Config) error {
	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return err
	}

	opt := &Options{
		MaxOpenConns:    20,
		MaxIdleConns:    10,
		ConnMaxLifetime: 10 * time.Minute,
		ConnMaxIdleTime: 5 * time.Minute,
	}
	if cfg.Options != nil {
		opt.rebuild(cfg.Options)
	}

	db.SetMaxOpenConns(opt.MaxOpenConns)
	db.SetMaxIdleConns(opt.MaxIdleConns)
	db.SetConnMaxLifetime(opt.ConnMaxLifetime)
	db.SetConnMaxIdleTime(opt.ConnMaxIdleTime)

	cli = ent.NewClient(
		ent.Driver(dialect.DebugWithContext(
			entsql.OpenDB(cfg.Driver, db),
			func(ctx context.Context, v ...any) {
				if config.ENV.Debug {
					log.Info(ctx, "SQL info", zap.String("SQL", fmt.Sprint(v...)))
				}
			}),
		),
	)

	return nil
}

// Client 返回Ent实例
func Client() *ent.Client {
	return cli
}
