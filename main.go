package main

import (
	"net/http"

	"github.com/felipeoli7eira/go-events-rest-api/db"
	"github.com/felipeoli7eira/go-events-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	db.Bootstrap()

	app.GET("/events", listEventsHandler)
	app.POST("/event", saveEventHandler)

	app.Run(":8080")
}

func listEventsHandler(gc *gin.Context) {
	list, err := models.ListEvents()

	if err != nil {
		gc.JSON(http.StatusInternalServerError, map[string]string {
			"tecnical_error": err.Error(),
			"error_message": "Failed to list events",
		})
		return
	}

	gc.JSON(http.StatusOK, map[string][]models.Event {
		"events": list,
	})
}

func saveEventHandler(gc *gin.Context) {
	var event models.Event

	var err error

	err = gc.ShouldBindJSON(&event)

	if err != nil {
		gc.JSON(http.StatusBadRequest, map[string]string {
			"tecnical_error": err.Error(),
			"error_message": "Invalid request body",
		})
		return
	}

	err = event.SaveEvent()

	if err != nil {
		gc.JSON(http.StatusInternalServerError, map[string]string {
			"tecnical_error": err.Error(),
			"error_message": "Failed to save event",
		})
		return
	}

	gc.JSON(http.StatusCreated, map[string]string {
		"message": "Event saved successfully",
	})
}
