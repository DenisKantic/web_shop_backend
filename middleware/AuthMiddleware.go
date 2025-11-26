package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web_shop_fis_backend/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// try to get the cookie
		tokenStr, err := c.Cookie("auth_token")
		if err != nil {
			fmt.Println("ERROR FIRST", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "Authentication token required"})
			return
		}

		// validate token
		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			c.SetCookie("auth_token", "", -1, "/", "", false, true) // maxAge -1 deletes cookie
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
