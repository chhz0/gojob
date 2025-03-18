package errcode

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Caused  string `json:"cause"`
	Message string `json:"msg"`
}

func New(code int, cause string, format string, args ...any) *Error {
	return &Error{
		Code:    code,
		Caused:  cause,
		Message: fmt.Sprintf(format, args...),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, cause: %s, msg: %s", e.Code, e.Caused, e.Message)
}

func (e *Error) WithMessage(format string, args ...any) *Error {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

func From(err error) *Error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); ok {
		return e
	}

	return ErrInternal
}
