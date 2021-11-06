package controllers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"task-managment/api/application"
	"task-managment/api/models"
	"task-managment/api/services"
	"task-managment/api/utils"
	customErrors "task-managment/api/utils/errors"
)

type userController struct {
	*application.BaseController
	userService *services.UserService
}

func NewUserController() *userController {
	return &userController{
		application.NewBaseController(),
		services.NewUserService(),
	}
}

func (userController *userController) Create(w http.ResponseWriter, r *http.Request) {

	userRequest := models.NewUser()
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		application.GetLogger().Error(err.Error())
		userController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}

	err := userRequest.Format().Sanitize().Validate()
	if err != nil {
		userController.JsonValidationErrors(w, err)
		return
	}

	userController.userService.User.Email = userRequest.Email
	foundUser, _ := userController.userService.Get()
	if foundUser != nil {
		userController.JsonValidationErrors(w, utils.ConvertToValidationError(errors.New(utils.EmailAlreadyExisted), "email"))
		return
	}

	userController.userService.User = userRequest
	response, err := userController.userService.Create()
	if err != nil {
		userController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}

	userController.Json(w, response, http.StatusOK)
}

func (userController *userController) List(w http.ResponseWriter, r *http.Request) {

	list, err := userController.userService.List(r)
	if err != nil {
		userController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}

	userController.Json(w, list, http.StatusOK)
}

func (userController *userController) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id = utils.SanitizeString(id)
	parsedId, _ := strconv.Atoi(id)
	userController.userService.User.Id = parsedId
	user, err := userController.userService.Get()
	if err != nil {
		userController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}
	userController.Json(w, user, http.StatusOK)
}

func (userController *userController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id = utils.SanitizeString(id)
	parsedId, _ := strconv.Atoi(id)
	userController.userService.User.Id = parsedId
	user, err := userController.userService.Delete()
	if err != nil {
		userController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}
	userController.Json(w, user, http.StatusNoContent)
}
