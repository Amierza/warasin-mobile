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
			// Psycholog
			routes.GET("/get-detail-psycholog", masterHandler.GetDetailPsycholog)

			// Practice
			routes.POST("/create-practice", psychologHandler.CreatePractice)
			routes.GET("/get-all-practice", psychologHandler.GetAllPractice)
			routes.PATCH("/update-practice/:id", psychologHandler.UpdatePractice)
			routes.DELETE("/delete-practice/:id", psychologHandler.DeletePractice)

			// Available Slot
			routes.GET("/get-all-available-slot", psychologHandler.GetAllAvailableSlot)

			// Consultation
			routes.GET("/get-all-consultation", psychologHandler.GetAllConsultation)
			routes.PATCH("/update-consultation/:id", psychologHandler.UpdateConsultation)
		}
	}
}
