package result

func OK(options ...ResultOption) Result {
	return New(0, "OK", options...)
}

func ErrParams(options ...ResultOption) Result {
	return New(10000, "params error", options...)
}

func ErrAuth(options ...ResultOption) Result {
	return New(20000, "unauthorized", options...)
}

func ErrPerm(options ...ResultOption) Result {
	return New(30000, "permission denied", options...)
}

func ErrNotFound(options ...ResultOption) Result {
	return New(40000, "entity not found", options...)
}

func ErrSystem(options ...ResultOption) Result {
	return New(50000, "internal server error", options...)
}

func ErrService(options ...ResultOption) Result {
	return New(60000, "internal service error", options...)
}
