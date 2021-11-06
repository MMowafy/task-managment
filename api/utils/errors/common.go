package payment_errors

import (
	"net/http"
	"task-managment/api/utils"
)

func GetHttpStatusCode(err error) int {
	switch err.(type) {

	case *InternalServerError:
		return http.StatusInternalServerError

	case *ValidationError:
		return http.StatusBadRequest

	case *NotFoundError:
		return http.StatusNotFound

	default:
		return http.StatusBadRequest

	}

}

func GetErrorStatusCode(err error) int {
	switch err.(type) {

	case *InternalServerError:
		internalServerError := err.(*InternalServerError)
		return internalServerError.GetErrorStatusCode()

	case *NotFoundError:
		notFoundError := err.(*NotFoundError)
		return notFoundError.GetErrorStatusCode()

	case *BadRequestError:
		badRequestError := err.(*BadRequestError)
		return badRequestError.GetErrorStatusCode()

	default:
		return utils.ErrorWithUnknownSystemCode
	}

}
