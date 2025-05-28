package main

import (
	"github.com/Aman5681/auth-micro/db"
	"github.com/Aman5681/auth-micro/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
