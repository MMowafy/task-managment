package services

import (
	"net/http"
	"reflect"
	"task-managment/api/infrastructure"
	"task-managment/api/models"
	"task-managment/api/utils"
	customErrors "task-managment/api/utils/errors"
)

type UserService struct {
	repository *infrastructure.UserRepository
	User       *models.User
}

func NewUserService() *UserService {
	return &UserService{
		infrastructure.NewUserRepository(),
		models.NewUser(),
	}
}

func (userService *UserService) Create() (*models.User, error) {
	createdUser := userService.repository.Create(userService.User)
	if createdUser == nil {
		return nil, customErrors.NewInternalServerError(utils.ErrorCreateUser, utils.ErrorCreateUserSystemCode)
	}
	return createdUser, nil
}

func (userService *UserService) Get() (*models.User, error) {
	foundUser := userService.repository.Find(userService.User)
	if foundUser == nil {
		return nil, customErrors.NewNotFoundError(utils.ErrorUserNotFound, utils.ErrorUserNotFoundSystemCode)
	}
	return foundUser, nil
}

func (userService *UserService) Delete() (*models.User, error) {
	err := userService.repository.Delete(userService.User)
	if err != nil {
		return nil, customErrors.NewNotFoundError(utils.ErrorUserNotFound, utils.ErrorUserNotFoundSystemCode)
	}
	return userService.User, nil
}

func (userService *UserService) List(r *http.Request) ([]models.User, error) {

	listRequest := utils.NewListRequest(r, reflect.ValueOf(userService.User).Elem())
	userList := userService.repository.List(listRequest)
	if userList == nil {
		return nil, customErrors.NewInternalServerError(utils.ErrorGeneralInternalError, utils.ErrorGeneralInternalErrorSystemCode)
	}
	return userList, nil
}
