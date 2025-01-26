package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/felipeoli7eira/go-events-rest-api/db"
	"github.com/felipeoli7eira/go-events-rest-api/models"
)

func main() {
	app := gin.Default()

	db.Bootstrap()

	app.GET("/events", listEventsHandler)
	app.GET("/event/:id", getEvent)
	app.POST("/event", saveEventHandler)

	app.Run(":8080")
}

func getEvent(gc *gin.Context) {
	id, err := strconv.ParseInt(gc.Param("id"), 10, 64)

	if err != nil {
		gc.JSON(http.StatusBadRequest, map[string]string {
			"error_message": "Invalid id",
			"tecnical_error": err.Error(),
		})
		return
	}

	event, err := models.GetEvent(id)

	if err != nil {
		gc.JSON(http.StatusInternalServerError, map[string]string {
			"error_message": "Could not fetch event",
			"tecnical_error": err.Error(),
		})
		return
	}

	gc.JSON(http.StatusInternalServerError, map[string]*models.Event {
		"data": event,
	})
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
