package log

import (
	"github.com/shenghui0779/yiigo/logger"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var log = logger.Debug()

// Init 初始化日志实例(如有多个实例，在此方法中初始化)
func Init() {
	cfg := &logger.Config{
		Filename: viper.GetString("log.path"),
	}

	opts := viper.GetStringMap("log.options")
	if len(opts) != 0 {
		cfg.Options = &logger.Options{
			MaxSize:    cast.ToInt(opts["max_size"]),
			MaxAge:     cast.ToInt(opts["max_age"]),
			MaxBackups: cast.ToInt(opts["max_backups"]),
			Compress:   cast.ToBool(opts["compress"]),
			Stderr:     cast.ToBool(opts["stderr"]),
		}
	}

	log = logger.New(cfg)
}
