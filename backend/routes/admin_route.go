package routes

import (
	"github.com/Amierza/warasin-mobile/backend/handler"
	"github.com/Amierza/warasin-mobile/backend/middleware"
	"github.com/Amierza/warasin-mobile/backend/service"
	"github.com/gin-gonic/gin"
)

func Admin(route *gin.Engine, adminHandler handler.IAdminHandler, masterHandler handler.IMasterHandler, jwtService service.IJWTService) {
	routes := route.Group("/api/v1/admin")
	{
		// Authentication
		routes.POST("/login", adminHandler.Login)
		routes.POST("/refresh-token", adminHandler.RefreshToken)

		routes.Use(middleware.Authentication(jwtService), middleware.RouteAccessControl(jwtService))
		{
			// Get All Role
			routes.GET("/get-all-role", adminHandler.GetAllRole)

			// CRUD User
			routes.POST("/create-user", adminHandler.CreateUser)
			routes.GET("/get-all-user", adminHandler.GetAllUser)
			routes.GET("/get-detail-user/:id", adminHandler.GetDetailUser)
			routes.PATCH("/update-user/:id", adminHandler.UpdateUser)
			routes.DELETE("/delete-user/:id", adminHandler.DeleteUser)

			// CRUD News
			routes.POST("/create-news", adminHandler.CreateNews)
			routes.GET("/get-all-news", adminHandler.GetAllNews)
			routes.GET("/get-detail-news/:id", adminHandler.GetDetailNews)
			routes.PATCH("/update-news/:id", adminHandler.UpdateNews)
			routes.DELETE("/delete-news/:id", adminHandler.DeleteNews)

			// CRUD Motivation Category
			routes.POST("/create-motivation-category", adminHandler.CreateMotivationCategory)
			routes.GET("/get-all-motivation-category", adminHandler.GetAllMotivationCategory)
			routes.GET("/get-detail-motivation-category/:id", adminHandler.GetDetailMotivationCategory)
			routes.PATCH("/update-motivation-category/:id", adminHandler.UpdateMotivationCategory)
			routes.DELETE("/delete-motivation-category/:id", adminHandler.DeleteMotivationCategory)

			// CRUD Motivation
			routes.POST("/create-motivation", adminHandler.CreateMotivation)
			routes.GET("/get-all-motivation", adminHandler.GetAllMotivation)
			routes.GET("/get-detail-motivation/:id", adminHandler.GetDetailMotivation)
			routes.PATCH("/update-motivation/:id", adminHandler.UpdateMotivation)
			routes.DELETE("/delete-motivation/:id", adminHandler.DeleteMotivation)

			// CRUD Psycholog
			routes.POST("/create-psycholog", adminHandler.CreatePsycholog)
			routes.GET("/get-all-psycholog", adminHandler.GetAllPsycholog)
			routes.GET("/get-detail-psycholog/:id", masterHandler.GetDetailPsycholog)
			routes.PATCH("/update-psycholog/:id", masterHandler.UpdatePsycholog)
			routes.DELETE("/delete-psycholog/:id", adminHandler.DeletePsycholog)

			// User Motivation
			routes.GET("/get-all-user-motivation", adminHandler.GetAllUserMotivation)

			// User Motivation
			routes.GET("/get-all-user-news", adminHandler.GetAllUserNews)

			// Consultation
			routes.GET("/get-all-consultation", adminHandler.GetAllConsultation)
		}
	}
}
