package main

import (
	"github.com/deuterium-cloud/go-web-crud/initializers"
	"github.com/deuterium-cloud/go-web-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Atom{})
}
