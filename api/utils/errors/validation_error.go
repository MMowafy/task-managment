package payment_errors

type ValidationError struct {
	Err error
}

func (validationError *ValidationError) Error() string {
	return validationError.Err.Error()
}

func NewValidationError(err error) *ValidationError {

	return &ValidationError{
		Err: err,
	}

}
