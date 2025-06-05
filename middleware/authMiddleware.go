package middleware

import (
	"gorm/utils" // replace with your actual import path
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")

		if authHeader == ""{
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer"{
			c.JSON(401, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		token := tokenParts[1]

		userID, err := utils.VerifyJWT(token)

		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}