package services

import (
	"net/http"
	"reflect"
	"task-managment/api/infrastructure"
	"task-managment/api/models"
	"task-managment/api/utils"
	customErrors "task-managment/api/utils/errors"
)

type NotificationService struct {
	repository   *infrastructure.NotificationRepository
	Notification *models.Notification
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		infrastructure.NewNotificationRepository(),
		models.NewNotification(),
	}
}

func (notificationService *NotificationService) CreateNotification(task *models.Task) (*models.Notification, error) {
	// TODO: as enhancment Prepare data and publish it to message queue to be consumed by mailer cli (another microservice)
	notificationService.Notification.UserId = task.UserId
	notificationService.Notification.TaskId = task.Id
	notificationService.Notification.Message = "New Task created for you"
	notificationService.Notification.IsRead = false
	createdNotification := notificationService.repository.Create(notificationService.Notification)
	if createdNotification == nil {
		return nil, customErrors.NewInternalServerError(utils.ErrorCreateNotification, utils.ErrorCreateNotificationSystemCode)
	}

	return createdNotification, nil
}

func (notificationService *NotificationService) Get() (*models.Notification, error) {
	foundNotification := notificationService.repository.Find(notificationService.Notification)
	if foundNotification == nil {
		return nil, customErrors.NewNotFoundError(utils.ErrorNotificationNotFound, utils.ErrorNotificationNotFoundSystemCode)
	}
	return foundNotification, nil
}

func (notificationService *NotificationService) List(r *http.Request) ([]models.Notification, error) {

	listRequest := utils.NewListRequest(r, reflect.ValueOf(notificationService.Notification).Elem())
	NotificationList := notificationService.repository.List(listRequest)
	if NotificationList == nil {
		return nil, customErrors.NewNotFoundError(utils.ErrorGeneralInternalError, utils.ErrorGeneralInternalErrorSystemCode)
	}
	return NotificationList, nil
}
