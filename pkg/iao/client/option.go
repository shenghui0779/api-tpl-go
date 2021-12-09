package client

// ActionOption configures how we set up the action
type ActionOption func(a *action)

// WithQuery specifies the `query` to Action.
func WithQuery(key, value string) ActionOption {
	return func(a *action) {
		a.query.Set(key, value)
	}
}

// WithBody specifies the `body` to Action.
func WithBody(f func() ([]byte, error)) ActionOption {
	return func(a *action) {
		a.body = f
	}
}

// WithDecode specifies the `decode` to Action.
func WithDecode(f func(resp []byte) error) ActionOption {
	return func(a *action) {
		a.decode = f
	}
}
