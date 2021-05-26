package result

type ResultOption func(r *response)

func WithMsg(msg string) ResultOption {
	return func(r *response) {
		r.msg = msg
	}
}

func WithErr(err error) ResultOption {
	return func(r *response) {
		r.err = err
	}
}

func WithData(data interface{}) ResultOption {
	return func(r *response) {
		r.data = data
	}
}
