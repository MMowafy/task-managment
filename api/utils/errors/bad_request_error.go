package payment_errors

type BadRequestError struct {
	Message    string
	StatusCode int
}

func (badRequestError *BadRequestError) Error() string {
	return badRequestError.Message
}

func (badRequestError *BadRequestError) GetErrorStatusCode() int {
	return badRequestError.StatusCode
}

func NewBadRequestError(message string, systemStatusCode int) *BadRequestError {
	return &BadRequestError{
		Message:    message,
		StatusCode: systemStatusCode}
}
