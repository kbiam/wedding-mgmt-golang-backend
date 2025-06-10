package api

import (
	"gorm/handlers"
	// "gorm/internal/auth"
	// "gorm/internal/guest"
	// "gorm/internal/twilio"
	"gorm/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine){

	// router.POST("/login", auth.Login)

	// guestRoutes :=  router.Group("/guests")
	// guestRoutes.Use(middleware.AuthMiddleware())
	// {
	// 	guestRoutes.POST("/", guest.AddGuest)
	// 	guestRoutes.GET("/", guest.GetGuests)
	// 	guestRoutes.PATCH("/:id/invite", guest.UpdateInvitationStatus)
	// 	guestRoutes.DELETE("/:id",guest.DeleteGuest)
	// 	guestRoutes.POST("/upload-excel",guest.UploadExcel)
	// }

	// protected := router.Group("/")
	// protected.Use(middleware.AuthMiddleware())
	// {


	// 	protected.POST("/send-template", twilio.SendMessage)
	// }
	router.POST("/login", handlers.Login)

	guestRoutes :=  router.Group("/guests")
	guestRoutes.Use(middleware.AuthMiddleware())
	{
		guestRoutes.POST("/", handlers.AddGuest)
		guestRoutes.GET("/", handlers.GetGuests)
		guestRoutes.PATCH("/:id/invite", handlers.UpdateInvitationStatus)
		guestRoutes.DELETE("/:id",handlers.DeleteGuest)
		guestRoutes.POST("/upload-excel",handlers.AddGuestsFromExcel)
	}

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{


		protected.POST("/send-template", handlers.SendMessage)
	}
}