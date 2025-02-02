package routes

import 	"github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.Engine) {
	app.GET("/events", listEventsHandler)
	app.GET("/event/:id", getEventHandler)
	app.POST("/event", saveEventHandler)
	app.PUT("/event/:id", updateEventHandler)
	app.DELETE("/event/:id", deleteEventHandler)

	app.POST("/signup", signUp)
	app.POST("/login", logIn)
}
