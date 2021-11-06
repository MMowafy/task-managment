package routes

import "task-managment/api/application"

func GetRoutes() []application.Route {
	var appRoutes []application.Route
	appRoutes = append(appRoutes, userRoutes...)
	appRoutes = append(appRoutes, taskRoutes...)
	appRoutes = append(appRoutes, notificationRoutes...)
	return appRoutes
}
