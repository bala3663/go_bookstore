package main

import (
	"Final-Project-gin/database"
	"Final-Project-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	port := "8000"
	router := gin.New() //router initialization
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.Book_keeperRoutes(router)
	routes.UserRoutes(router)
	router.Run(":" + port)

}
