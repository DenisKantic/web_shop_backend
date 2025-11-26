package Account

import (
	"domko_backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAuthLogin(c *gin.Context) {
	tokenCookie, err := c.Cookie("auth_token")
	fmt.Println("TOKEN GETTING", tokenCookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "Authentication token missing"})
		fmt.Println("ERROR", err)
		return
	}

	claims, validateErr := utils.ValidateToken(tokenCookie)

	if validateErr != nil {
		c.SetCookie("auth_token", "", -1, "/", "", false, true) // clear the invalid cookie to prevent repeated failures
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login"})
		fmt.Println("ERROR", validateErr)
		return
	}

	c.Set("claims", claims)

	fmt.Println("USER ROLE", claims.Role)
	c.JSON(http.StatusOK, gin.H{"role": claims.Role, "full_name": claims.FullName})
}
