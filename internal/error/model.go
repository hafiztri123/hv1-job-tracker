package appError

type AppError struct {
	Err        error  `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

func (e *AppError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(err error, message string, statusCode int) *AppError {
	return &AppError{
		Err:        err,
		Message:    message,
		StatusCode: statusCode,
	}
}
