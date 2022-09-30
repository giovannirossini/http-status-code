package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/giovannirossini/http-status-code/controllers"
)

func main() {
	port := os.Getenv("PORT")
	router := gin.Default()
	router.GET("/", controllers.GetStatus)
	router.GET("/:id", controllers.GetStatusById)

	router.Run(":" + port)
}
