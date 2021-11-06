package controllers

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"task-managment/api/application"
	"task-managment/api/services"
	"task-managment/api/utils"
	customErrors "task-managment/api/utils/errors"
)

type NotificationController struct {
	*application.BaseController
	NotificationService *services.NotificationService
}

func NewNotificationController() *NotificationController {
	return &NotificationController{
		application.NewBaseController(),
		services.NewNotificationService(),
	}
}

func (NotificationController *NotificationController) List(w http.ResponseWriter, r *http.Request) {

	list, err := NotificationController.NotificationService.List(r)
	if err != nil {
		NotificationController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}

	NotificationController.Json(w, list, http.StatusOK)
}

func (NotificationController *NotificationController) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id = utils.SanitizeString(id)
	parsedId, _ := strconv.Atoi(id)
	NotificationController.NotificationService.Notification.Id = parsedId
	Notification, err := NotificationController.NotificationService.Get()
	if err != nil {
		NotificationController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}
	NotificationController.Json(w, Notification, http.StatusOK)
}
