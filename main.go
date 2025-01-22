package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/events", listEvents)

	app.Run(":8080")
}

func listEvents(c *gin.Context) {
	c.JSON(http.StatusOK, map[string] any {
		"status": "OK",
	})
}
