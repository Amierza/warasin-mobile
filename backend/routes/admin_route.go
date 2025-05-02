package routes

import (
	"github.com/Amierza/warasin-mobile/backend/handler"
	"github.com/Amierza/warasin-mobile/backend/service"
	"github.com/gin-gonic/gin"
)

func Admin(route *gin.Engine, adminHandler handler.IAdminHandler, jwtService service.IJWTService) {
	routes := route.Group("/api/v1/admin")
	{
		routes.POST("/login", adminHandler.Login)
		routes.POST("/refresh-token", adminHandler.RefreshToken)
	}
}
