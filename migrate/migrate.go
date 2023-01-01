package main

import (
	"gitlab.com/deuterium/web-crud/initializers"
	"gitlab.com/deuterium/web-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Atom{})
}
