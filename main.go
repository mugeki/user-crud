package main

import (
	"user-crud/config"
	"user-crud/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()
	e := routes.NewRoutes()
	e.Start(":8000")
}