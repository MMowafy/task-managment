package main

import (
	_ "github.com/lib/pq"
	"task-managment/api/application"
	"task-managment/api/routes"
)

func main() {
	app := application.NewApplication()
	app.MigrateAndSeedDB()
	app.SetRoutes(routes.GetRoutes())
	app.StartServer()
}
