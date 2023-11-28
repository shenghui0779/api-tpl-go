package value

// EmptyEncMode 值为空时的Encode模式
type EmptyEncMode int

const (
	EmptyEncDefault EmptyEncMode = iota // 默认：bar=baz&foo=
	EmptyEncIgnore                      // 忽略：bar=baz
	EmptyEncOnlyKey                     // 仅保留Key：bar=baz&foo
)

type options struct {
	escape     bool
	emptyMode  EmptyEncMode
	ignoreKeys map[string]struct{}
}

// Option V Encode 选项
type Option func(o *options)

// WithEmptyEncMode 设置值为空时的Encode模式
func WithEmptyEncMode(mode EmptyEncMode) Option {
	return func(o *options) {
		o.emptyMode = mode
	}
}

// WithKVEscape 设置K-V是否需要QueryEscape
func WithKVEscape() Option {
	return func(o *options) {
		o.escape = true
	}
}

// WithIgnoreKeys 设置Encode时忽略的key
func WithIgnoreKeys(keys ...string) Option {
	return func(o *options) {
		for _, k := range keys {
			o.ignoreKeys[k] = struct{}{}
		}
	}
}
