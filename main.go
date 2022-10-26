package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mongo/connection/configs"
	"github.com/mongo/connection/routes"
)

func main() {

	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)

	router.Run("localhost:8080")
}
