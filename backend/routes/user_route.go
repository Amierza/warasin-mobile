package routes

import (
	"github.com/Amierza/warasin-mobile/backend/handler"
	"github.com/Amierza/warasin-mobile/backend/middleware"
	"github.com/Amierza/warasin-mobile/backend/service"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userHandler handler.IUserHandler, jwtService service.IJWTService) {
	routes := route.Group("/api/v1/user")
	{
		routes.POST("/register", userHandler.Register)
		routes.POST("/login", userHandler.Login)

		routes.POST("/send-forgot-password-email", userHandler.SendForgotPasswordEmail)
		routes.GET("/forgot-password", userHandler.ForgotPassword)
		routes.POST("/update-password", userHandler.UpdatePassword)

		routes.POST("/send-verification-email", userHandler.SendVerificationEmail)
		routes.GET("/verify-email", userHandler.VerifyEmail)

		routes.GET("/get-detail-user", middleware.Authentication(jwtService), userHandler.GetDetailUser)
	}
}
