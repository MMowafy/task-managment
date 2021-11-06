package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"task-managment/api/application"
	"task-managment/api/models"
	"task-managment/api/services"
	"task-managment/api/utils"
	customErrors "task-managment/api/utils/errors"
)

type TaskController struct {
	*application.BaseController
	taskService *services.TaskService
}

func NewTaskController() *TaskController {
	return &TaskController{
		application.NewBaseController(),
		services.NewTaskService(),
	}
}
func (taskController *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	taskRequest := models.NewTask()
	if err := json.NewDecoder(r.Body).Decode(&taskRequest); err != nil {
		application.GetLogger().Error(err.Error())
		taskController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}

	err := taskRequest.Format().Sanitize().Validate()
	if err != nil {
		taskController.JsonValidationErrors(w, err)
		return
	}
	taskController.taskService.Task = taskRequest
	response, err := taskController.taskService.Create()
	if err != nil {
		taskController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}

	taskController.Json(w, response, http.StatusOK)
}

func (taskController *TaskController) List(w http.ResponseWriter, r *http.Request) {

	list, err := taskController.taskService.List(r)
	if err != nil {
		taskController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}

	taskController.Json(w, list, http.StatusOK)
}

func (taskController *TaskController) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id = utils.SanitizeString(id)
	parsedId, _ := strconv.Atoi(id)
	taskController.taskService.Task.Id = parsedId
	task, err := taskController.taskService.Get()
	if err != nil {
		taskController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}
	taskController.Json(w, task, http.StatusOK)
}

func (taskController *TaskController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id = utils.SanitizeString(id)
	parsedId, _ := strconv.Atoi(id)
	taskController.taskService.Task.Id = parsedId
	task, err := taskController.taskService.Delete()
	if err != nil {
		taskController.JsonError(w, err.Error(), customErrors.GetHttpStatusCode(err), customErrors.GetErrorStatusCode(err))
		return
	}
	taskController.Json(w, task, http.StatusNoContent)
}
