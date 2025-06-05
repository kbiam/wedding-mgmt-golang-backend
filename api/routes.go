package api

import (
	"github.com/gin-gonic/gin"
	"gorm/internal/guest"
	"gorm/internal/auth"
	"gorm/internal/twilio"
	"gorm/middleware"
)

func SetupRoutes(router *gin.Engine){

	router.POST("/login", auth.Login)

	guestRoutes :=  router.Group("/guests")
	guestRoutes.Use(middleware.AuthMiddleware())
	{
		guestRoutes.POST("/", guest.AddGuest)
		guestRoutes.GET("/", guest.GetGuests)
		guestRoutes.PATCH("/:id/invite", guest.UpdateInvitationStatus)
		guestRoutes.DELETE("/:id",guest.DeleteGuest)
	}

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{


		protected.POST("/send-template", twilio.SendMessage)
	}
}