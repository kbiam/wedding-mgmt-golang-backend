package auth

import (
	
	"github.com/gin-gonic/gin"
	"gorm/db"
	"gorm/utils"
)

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context){
	var user AdminUser
	var loginData LoginBody

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	if err := db.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}

	if user.Password != loginData.Password {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	
	jwtToken, err := utils.GenerateJWT(user.ID); if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate JWT token"})
		return
	}


	c.JSON(200, gin.H{"message": "Login successful", "user": user, "token": jwtToken})
}
