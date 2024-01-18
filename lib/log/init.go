package log

import (
	"github.com/shenghui0779/yiigo/logger"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var log = logger.Debug()

// Init 初始化日志实例(如有多个实例，在此方法中初始化)
func Init() {
	log = logger.New(buildCfg(viper.GetString("log.path"), viper.GetStringMap("log.options")))
}

func buildCfg(path string, opts map[string]any) *logger.Config {
	cfg := &logger.Config{
		Filename: path,
	}

	if len(opts) != 0 {
		cfg.Options = &logger.Options{
			MaxSize:    cast.ToInt(opts["max_size"]),
			MaxAge:     cast.ToInt(opts["max_age"]),
			MaxBackups: cast.ToInt(opts["max_backups"]),
			Compress:   cast.ToBool(opts["compress"]),
			Stderr:     cast.ToBool(opts["stderr"]),
		}
	}

	return cfg
}
