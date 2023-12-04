package log

import (
	"os"
	"time"

	"github.com/spf13/cast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config 日志初始化配置
type Config struct {
	// Filename 日志名称
	Filename string `json:"filename"`
	// Options 日志选项
	Options *Options `json:"options"`
}

// Options 日志配置选项
type Options struct {
	// MaxSize 当前文件多大时轮替；默认：100MB
	MaxSize int `json:"max_size"`
	// MaxAge 轮替的旧文件最大保留时长；默认：不限
	MaxAge int `json:"max_age"`
	// MaxBackups 轮替的旧文件最大保留数量；默认：不限
	MaxBackups int `json:"max_backups"`
	// Compress 轮替的旧文件是否压缩；默认：不压缩
	Compress bool `json:"compress"`
	// Stderr 是否输出到控制台
	Stderr bool `json:"stderr"`
	// ZapOptions Zap日志选项
	ZapOptions []zap.Option `json:"zap_options"`
}

func debug(options ...zap.Option) *zap.Logger {
	cfg := zap.NewDevelopmentConfig()

	cfg.DisableCaller = true
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = MyTimeEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	l, _ := cfg.Build(options...)

	return l
}

func New(cfg *Config) *zap.Logger {
	if len(cfg.Filename) == 0 {
		return debug(cfg.Options.ZapOptions...)
	}

	c := zap.NewProductionEncoderConfig()
	c.TimeKey = "time"
	c.EncodeTime = MyTimeEncoder
	c.EncodeCaller = zapcore.FullCallerEncoder

	w := &lumberjack.Logger{
		Filename:  cfg.Filename,
		LocalTime: true,
	}
	ws := make([]zapcore.WriteSyncer, 0, 2)
	if cfg.Options != nil {
		w.MaxSize = cfg.Options.MaxSize
		w.MaxAge = cfg.Options.MaxAge
		w.MaxBackups = cfg.Options.MaxBackups
		w.Compress = cfg.Options.Compress

		if cfg.Options.Stderr {
			ws = append(ws, zapcore.Lock(os.Stderr))
		}
	}
	ws = append(ws, zapcore.AddSync(w))

	core := zapcore.NewCore(zapcore.NewJSONEncoder(c), zapcore.NewMultiWriteSyncer(ws...), zap.DebugLevel)

	return zap.New(core, cfg.Options.ZapOptions...)
}

// MyTimeEncoder 自定义时间格式化
func MyTimeEncoder(t time.Time, e zapcore.PrimitiveArrayEncoder) {
	e.AppendString(t.In(time.FixedZone("CST", 8*3600)).Format(time.DateTime))
}

func buildCfg(filename string, options map[string]any) *Config {
	cfg := &Config{
		Filename: filename,
	}

	if len(options) != 0 {
		cfg.Options = &Options{
			MaxSize:    cast.ToInt("max_size"),
			MaxAge:     cast.ToInt("log.max_age"),
			MaxBackups: cast.ToInt("log.max_backups"),
			Compress:   cast.ToBool("log.compress"),
			Stderr:     cast.ToBool("log.stderr"),
		}
	}

	return cfg
}
