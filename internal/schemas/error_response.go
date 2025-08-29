package schemas

type ErrorStruc struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *ErrorStruc) Error() string {
	return e.Message
}

func ErrorResponse(code int, message string, err error) *ErrorStruc {
	return &ErrorStruc{
		StatusCode: code,
		Message:    message,
		Err:        err,
	}
}
