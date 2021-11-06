package payment_errors

type NotFoundError struct {
	Message    string
	StatusCode int
}

func (notFoundError *NotFoundError) Error() string {
	return notFoundError.Message
}

func (notFoundError *NotFoundError) GetErrorStatusCode() int {
	return notFoundError.StatusCode
}

func NewNotFoundError(message string, systemStatusCode int) *NotFoundError {
	return &NotFoundError{
		Message:    message,
		StatusCode: systemStatusCode}
}
