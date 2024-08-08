package routes

import (
	"github.com/Maged-Zaki/gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(server *gin.Engine, path string) {
	routerGroup := server.Group(path)

	routerGroup.POST("/signup", controllers.Signup)
	routerGroup.POST("/login", controllers.Login)
}
