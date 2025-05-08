package routes

import (
	"github.com/Amierza/warasin-mobile/backend/handler"
	"github.com/Amierza/warasin-mobile/backend/middleware"
	"github.com/Amierza/warasin-mobile/backend/service"
	"github.com/gin-gonic/gin"
)

func Admin(route *gin.Engine, adminHandler handler.IAdminHandler, jwtService service.IJWTService) {
	routes := route.Group("/api/v1/admin")
	{
		routes.POST("/login", adminHandler.Login)
		routes.POST("/refresh-token", adminHandler.RefreshToken)
		routes.Use(middleware.Authentication(jwtService), middleware.RouteAccessControl(jwtService))
		{
			// CRUD User
			routes.POST("/create-user", adminHandler.CreateUser)
			routes.GET("/get-all-user", adminHandler.GetAllUser)
			routes.PATCH("/update-user/:id", adminHandler.UpdateUser)
			routes.DELETE("/delete-user/:id", adminHandler.DeleteUser)

			// CRUD News
			routes.POST("/create-news", adminHandler.CreateNews)
			routes.GET("/get-all-news", adminHandler.GetAllNews)
			routes.PATCH("/update-news/:id", adminHandler.UpdateNews)
			routes.DELETE("/delete-news/:id", adminHandler.DeleteNews)

			// CRUD Motivation Category
			routes.POST("/create-motivation-category", adminHandler.CreateMotivationCategory)
			routes.GET("/get-all-motivation-category", adminHandler.GetAllMotivationCategory)
			routes.PATCH("/update-motivation-category/:id", adminHandler.UpdateMotivationCategory)
			routes.DELETE("/delete-motivation-category/:id", adminHandler.DeleteMotivationCategory)

			// CRUD Motivation
			routes.POST("/create-motivation", adminHandler.CreateMotivation)
			routes.GET("/get-all-motivation", adminHandler.GetAllMotivation)
		}
	}
}
