package routes

import (
	"net/http"
	"task-managment/api/application"
	"task-managment/api/controllers"
)

const TaskResourceName = "/tasks"

var taskRoutes = []application.Route{
	{http.MethodPost, TaskResourceName,
		func(writer http.ResponseWriter, request *http.Request) {
			controllers.NewTaskController().Create(writer, request)
		},
	},
	{http.MethodGet, TaskResourceName,
		func(writer http.ResponseWriter, request *http.Request) {
			controllers.NewTaskController().List(writer, request)
		},
	},
	{http.MethodGet, TaskResourceName + "/{id}",
		func(writer http.ResponseWriter, request *http.Request) {
			controllers.NewTaskController().Get(writer, request)
		},
	},
	{http.MethodDelete, TaskResourceName + "/{id}",
		func(writer http.ResponseWriter, request *http.Request) {
			controllers.NewTaskController().Delete(writer, request)
		},
	},
}
