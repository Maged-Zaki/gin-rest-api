package routes

import (
	"github.com/Maged-Zaki/gin-rest-api/controllers"
	"github.com/Maged-Zaki/gin-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(server *gin.Engine, path string) {
	routerGroup := server.Group(path)

	routerGroup.Use(middlewares.ValidateJWT)
	routerGroup.DELETE("/users/:id", controllers.DeleteUser)

}
