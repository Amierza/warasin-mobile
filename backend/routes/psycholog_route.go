package routes

import (
	"github.com/Amierza/warasin-mobile/backend/handler"
	"github.com/Amierza/warasin-mobile/backend/middleware"
	"github.com/Amierza/warasin-mobile/backend/service"
	"github.com/gin-gonic/gin"
)

func Psycholog(route *gin.Engine, psychologHandler handler.IPsychologHandler, masterHandler handler.IMasterHandler, jwtService service.IJWTService) {
	routes := route.Group("/api/v1/psycholog")
	{
		routes.POST("/login", psychologHandler.Login)
		routes.POST("/refresh-token", psychologHandler.RefreshToken)
		routes.Use(middleware.Authentication(jwtService), middleware.RouteAccessControl(jwtService))
		{
			routes.GET("/get-detail-psycholog", masterHandler.GetDetailPsycholog)
		}
	}
}
