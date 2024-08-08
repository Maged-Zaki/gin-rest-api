package routes

import (
	"github.com/Maged-Zaki/gin-rest-api/controllers"
	"github.com/Maged-Zaki/gin-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func EventsRoutes(server *gin.Engine, path string) {
	routerGroup := server.Group(path)

	routerGroup.Use(middlewares.ValidateJWT)
	routerGroup.POST("", controllers.CreateEvent)
	routerGroup.PUT("/:id", controllers.UpdateEvent)
	routerGroup.DELETE("/:id", controllers.DeleteEvent)
	routerGroup.GET("/:id", controllers.GetEvent)
	routerGroup.GET("", controllers.GetAllEvents)
}
