package main

import (
	"golangrestapi/db"
	"golangrestapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitDb()
	routes.RegisterRoutes(server)
	server.Run()
}
