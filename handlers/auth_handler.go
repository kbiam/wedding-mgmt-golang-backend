package handlers

import (
	auth_request "gorm/models/request/auth"
	"gorm/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	var loginData auth_request.LoginBody

	 if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	 }
	user, token, err := services.Login(loginData)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
		"token":   token,
	})
}