package routes

import (
	"net/http"
	"task-managment/api/application"
	"task-managment/api/controllers"
)

const UserResourceName = "/users"

var userRoutes = []application.Route{
	{http.MethodPost, UserResourceName,
		func(writer http.ResponseWriter, request *http.Request) {
			controllers.NewUserController().Create(writer, request)
		},
	},
	{http.MethodGet, UserResourceName,
		func(writer http.ResponseWriter, request *http.Request) {
			controllers.NewUserController().List(writer, request)
		},
	},
	{http.MethodGet, UserResourceName + "/{id}",
		func(writer http.ResponseWriter, request *http.Request) {
			controllers.NewUserController().Get(writer, request)
		},
	},
	{http.MethodDelete, UserResourceName + "/{id}",
		func(writer http.ResponseWriter, request *http.Request) {
			controllers.NewUserController().Delete(writer, request)
		},
	},
}
