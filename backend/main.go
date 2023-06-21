package main

import (
	"backend/configs"
	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDatabase()

	router := gin.Default()
	v1 := router.Group("/v1")
	routes.AddRoutes(v1)
	router.Use(cors.Default())

	router.Run(":8877")
}
