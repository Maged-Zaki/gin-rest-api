package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Auth
	AuthRoutes(server, "/auth")
	// Users
	UsersRoutes(server, "/users")
	// Events
	EventsRoutes(server, "/events")
}
