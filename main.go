package main

import (
	"github.com/gin-gonic/gin"

	"github.com/felipeoli7eira/go-events-rest-api/db"
	"github.com/felipeoli7eira/go-events-rest-api/routes"
)

func main() {
	app := gin.Default()

	db.Bootstrap()

	routes.RegisterRoutes(app)

	app.Run(":8080")
}
