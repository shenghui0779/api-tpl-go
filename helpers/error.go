package helpers

type Status interface {
	error
	Code() Code
}

type Error struct {
	code Code
	msg  string
}

func (e *Error) Code() Code {
	return e.code
}

func (e *Error) Error() string {
	return e.msg
}

// Err returns an error
func Err(code Code, msg ...string) error {
	err := &Error{code: code}

	if len(msg) != 0 {
		err.msg = msg[0]

		return err
	}

	err.msg = "Whoops! Something Wrong!"

	if m, ok := codeM[code]; ok {
		err.msg = m
	}

	return err
}

// ErrCode returns error code
func ErrCode(err error) Code {
	if v, ok := err.(Status); ok {
		return v.Code()
	}

	return ErrSystem
}

// ErrMsg returns error msg
func ErrMsg(err error) string {
	if v, ok := err.(Status); ok {
		return v.Error()
	}

	return codeM[ErrSystem]
}
