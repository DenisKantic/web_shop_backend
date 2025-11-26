package Account

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"web_shop_fis_backend/database"
	"web_shop_fis_backend/models"
	"web_shop_fis_backend/utils"
)

func ManualLogin(c *gin.Context) {

	var login models.ManualLoginUserRequest
	var user models.User
	// userIP := c.ClientIP()

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := database.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if !user.IsEmailVerified {
		c.JSON(http.StatusForbidden, gin.H{"error": "Your account is not verified."})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(login.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	fmt.Println("USER ROLE BEFORE SENDING TO JWT", user.Role)

	token, err := utils.GenerateJWT(user.ID, user.FullName, user.Role, "auth", 7*24*time.Hour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to login"})
		return
	}

	fmt.Println("AFTER ROLE", user.Role)
	fmt.Println("NEW LOGIN TOKEN", token)

	c.SetCookie("auth_token", token, 60*60*24, "/", "localhost", false, true)
	claims, validateErr := utils.ValidateToken(token)
	if validateErr != nil {
		c.SetCookie("auth_token", "", -1, "/", "localhost", false, true) // clear the invalid cookie to prevent repeated failures
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login"})
		return
	}

	fmt.Println("TOKEN AFTER VALIDATE", claims.UserID)

	c.JSON(http.StatusOK, gin.H{"message": claims})

}

func Logout(c *gin.Context) {
	c.SetCookie("auth_token", "", -1, "/", "", false, false)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
}
