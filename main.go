package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/deuterium/web-crud/controllers"
	"gitlab.com/deuterium/web-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/atoms/pagination", controllers.GetAtomsWithPagination)
	router.GET("/atoms", controllers.GetAtoms)
	router.GET("/atoms/:id", controllers.GetAtomById)
	router.POST("/atoms", controllers.SaveNewAtom)
	router.PUT("/atoms/:id", controllers.UpdateAtom)
	router.DELETE("/atoms/:id", controllers.DeleteAtom)
	router.Run() // listen and serve on 0.0.0.0:${PORT}
}
