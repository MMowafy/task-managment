package payment_errors

type InternalServerError struct {
	Message    string
	StatusCode int
}

func (internalServerError *InternalServerError) Error() string {
	return internalServerError.Message
}

func (internalServerError *InternalServerError) GetErrorStatusCode() int {
	return internalServerError.StatusCode
}

func NewInternalServerError(message string, systemStatusCode int) *InternalServerError {
	return &InternalServerError{
		Message:    message,
		StatusCode: systemStatusCode}
}
