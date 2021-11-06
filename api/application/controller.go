package application

import (
	"encoding/json"
	"net/http"
)

type BaseController struct{}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (baseController *BaseController) json(res http.ResponseWriter, payload interface{}, statusCode int) {
	response, _ := json.Marshal(payload)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	res.Write(response)
}

func (baseController *BaseController) Json(res http.ResponseWriter, payload interface{}, statusCode int) {
	baseController.json(res, payload, statusCode)
}

func (baseController *BaseController) JsonError(res http.ResponseWriter, msg string, httpStatusCode int, statusCode int, ) {

	response := map[string]interface{}{
		"systemStatusCode": statusCode,
		"status":           httpStatusCode,
		"message":          msg,
	}

	baseController.json(res, response, httpStatusCode)
}

func (baseController *BaseController) JsonValidationErrors(res http.ResponseWriter, err error) {
	response := map[string]interface{}{
		"systemStatusCode": 5000,
		"status":           http.StatusBadRequest,
		"message":          err.Error(),
		"errors":           err,
	}
	baseController.json(res, response, http.StatusBadRequest)
}
