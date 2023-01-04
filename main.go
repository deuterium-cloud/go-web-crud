package main

import (
	"github.com/deuterium-cloud/go-web-crud/controllers"
	"github.com/deuterium-cloud/go-web-crud/initializers"
	"github.com/gin-gonic/gin"
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

	router.MaxMultipartMemory = 8 << 20 // 8 Mb
	router.POST("/upload", controllers.UploadFile)
	router.Run() // listen and serve on 0.0.0.0:${PORT}
}
