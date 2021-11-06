package services

import (
	"net/http"
	"reflect"
	"task-managment/api/infrastructure"
	"task-managment/api/models"
	"task-managment/api/utils"
	customErrors "task-managment/api/utils/errors"
	"time"
)

type TaskService struct {
	repository          *infrastructure.TaskRepository
	Task                *models.Task
	NotificationService *NotificationService
	MailerService       *MailerService
}

func NewTaskService() *TaskService {
	return &TaskService{
		infrastructure.NewTaskRepository(),
		models.NewTask(),
		NewNotificationService(),
		NewMailerService(),
	}
}

func (taskService *TaskService) Create() (*models.Task, error) {
	taskService.Task.EndAt = taskService.Task.StartAt.Add(time.Duration(taskService.Task.Duration) * time.Hour)
	foundTask := taskService.repository.FindOverlappingTasks(taskService.Task)
	if foundTask != nil {
		return nil, customErrors.NewBadRequestError(utils.ErrorTaskOverlappingFound, utils.ErrorTaskOverlappingFoundSystemCode)
	}
	createdTask := taskService.repository.Create(taskService.Task)
	if createdTask == nil {
		return nil, customErrors.NewInternalServerError(utils.ErrorCreateTask, utils.ErrorCreateTaskSystemCode)
	}

	go taskService.NotificationService.CreateNotification(taskService.Task)
	go taskService.MailerService.SendEmail(taskService.Task)
	return createdTask, nil
}

func (taskService *TaskService) Get() (*models.Task, error) {
	foundTask := taskService.repository.Find(taskService.Task)
	if foundTask == nil {
		return nil, customErrors.NewNotFoundError(utils.ErrorTaskNotFound, utils.ErrorTaskNotFoundSystemCode)
	}
	return foundTask, nil
}

func (taskService *TaskService) Delete() (*models.Task, error) {
	err := taskService.repository.Delete(taskService.Task)
	if err != nil {
		return nil, customErrors.NewNotFoundError(utils.ErrorTaskNotFound, utils.ErrorTaskNotFoundSystemCode)
	}
	return taskService.Task, nil
}

func (taskService *TaskService) List(r *http.Request) ([]models.Task, error) {

	listRequest := utils.NewListRequest(r, reflect.ValueOf(taskService.Task).Elem())
	taskList := taskService.repository.List(listRequest)
	if taskList == nil {
		return nil, customErrors.NewNotFoundError(utils.ErrorGeneralInternalError, utils.ErrorGeneralInternalErrorSystemCode)
	}
	return taskList, nil
}
