package handlers

import (
	twilio_request "gorm/models/request/twilio"
	"gorm/services"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context){
	var body twilio_request.TwilioApiBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := services.SendMessage(body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Message sent successfully",
		"response": resp,
	})
}