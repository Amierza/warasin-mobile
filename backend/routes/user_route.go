package routes

import (
	"github.com/Amierza/warasin-mobile/backend/handler"
	"github.com/Amierza/warasin-mobile/backend/middleware"
	"github.com/Amierza/warasin-mobile/backend/service"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userHandler handler.IUserHandler, masterHandler handler.IMasterHandler, jwtService service.IJWTService) {
	routes := route.Group("/api/v1/user")
	{
		// Authentication
		routes.POST("/register", userHandler.Register)
		routes.POST("/login", userHandler.Login)
		routes.POST("/refresh-token", userHandler.RefreshToken)

		// Forgot Password
		routes.POST("/send-forgot-password-email", userHandler.SendForgotPasswordEmail)
		routes.GET("/forgot-password", userHandler.ForgotPassword)
		routes.PATCH("/update-password", userHandler.UpdatePassword)

		// Verification Email
		routes.POST("/send-verification-email", userHandler.SendVerificationEmail)
		routes.GET("/verify-email", userHandler.VerifyEmail)

		routes.Use(middleware.Authentication(jwtService), middleware.RouteAccessControl(jwtService))
		{
			// User
			routes.GET("/get-detail-user", userHandler.GetDetailUser)
			routes.PATCH("/update-user", userHandler.UpdateUser)

			// News
			routes.GET("/get-all-news", userHandler.GetAllNews)
			routes.GET("/get-detail-news/:id", userHandler.GetDetailNews)

			// Motivation
			routes.GET("/get-all-motivation", userHandler.GetAllMotivation)
			routes.GET("/get-detail-motivation/:id", userHandler.GetDetailMotivation)

			// Consultation
			routes.POST("create-consultation", userHandler.CreateConsultation)
			routes.GET("get-all-consultation", userHandler.GetAllConsultation)
			routes.GET("get-detail-consultation/:id", userHandler.GetDetailConsultation)
			routes.PATCH("update-consultation/:id", userHandler.UpdateConsultation)
			routes.DELETE("delete-consultation/:id", userHandler.DeleteConsultation)

			// Psycholog
			routes.GET("get-all-psycholog", userHandler.GetAllPsycholog)
			routes.GET("get-detail-psycholog/:id", userHandler.GetDetailPsycholog)

			// Practice
			routes.GET("/get-all-practice/:psyID", userHandler.GetAllPractice)

			// Available Slot
			routes.GET("/get-all-available-slot/:psyID", userHandler.GetAllAvailableSlot)

			// News Detail
			routes.POST("create-news-detail", userHandler.CreateNewsDetail)
			routes.GET("get-all-news-detail", userHandler.GetAllNewsDetail)

			// User Motivation
			routes.POST("create-user-motivation", userHandler.CreateUserMotivation)
			routes.GET("get-all-user-motivation", userHandler.GetAllUserMotivation)
		}
	}
}
