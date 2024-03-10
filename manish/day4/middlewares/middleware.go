package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"manish/day4/service"
)

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		fmt.Println("tokenString", tokenString)
		if err = service.VerifyToken(tokenString, role); err != nil {
			c.JSON(401, gin.H{
				"error":   "Invalid token",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
