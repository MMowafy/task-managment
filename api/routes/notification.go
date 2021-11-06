package routes

import (
	"net/http"
	"task-managment/api/application"
	"task-managment/api/controllers"
)

const NotificationResourceName = UserResourceName + "/{userId}/notifications"

var notificationRoutes = []application.Route{
	{http.MethodGet, NotificationResourceName,
		func(writer http.ResponseWriter, request *http.Request) {
			controllers.NewNotificationController().List(writer, request)
		},
	},
	{http.MethodGet, NotificationResourceName + "/{id}",
		func(writer http.ResponseWriter, request *http.Request) {
			controllers.NewNotificationController().Get(writer, request)
		},
	},
}
