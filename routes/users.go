package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/felipeoli7eira/go-events-rest-api/models"
)

func signUp(gc *gin.Context) {
	var user models.User

	err := gc.ShouldBindJSON(&user)

	if err != nil {
		gc.JSON(http.StatusBadRequest, map[string]string {
			"error_message": "Invalid request body",
			"tecnical_error": err.Error(),
		})
		return
	}

	err = user.Save()

	if err != nil {
		gc.JSON(http.StatusInternalServerError, map[string]string {
			"error_message": "Could not save user",
			"tecnical_error": err.Error(),
		})
		return
	}

	gc.JSON(http.StatusCreated, map[string]string {
		"message": "User created successfully",
	})
}
