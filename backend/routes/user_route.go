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

			// Consultation
			routes.POST("create-consultation", userHandler.CreateConsultation)
			routes.GET("get-all-consultation", userHandler.GetAllConsultation)
		}
	}
}
