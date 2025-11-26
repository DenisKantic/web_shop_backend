// Package Account Package controllers  package used for handling API methods
package Account

import (
	"domko_backend/controllers/EmailController"
	"domko_backend/database"
	"domko_backend/models"
	"domko_backend/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// RegisterWithEmailOwnerUser handles user registration with email and password*
//
// It performs the following steps:
//
//   - parses and validates the incoming JSON body with user model
//
//   - checks if the user already exists by email
//     -hashes the user's password securely
//     -stores the new user in the database
//
//   - returns appropriate HTTP status codes and messages
//
// Expected input JSON:
//
//	{
//		"email": "user@test.com",
//		"username": "John Doe",
//		"password: "yourPassword",
//		"location": "New York",
//	}
//
// Possible responses:
//
//   - 201 Created: user successfully registered
//
//   - 400 Bad request: validation failed or email in use
//
//   - 500 Internal Server Error: on DB or hashing failure
func RegisterWithEmailOwnerUser(c *gin.Context) {

	userIP := c.ClientIP()

	fmt.Println("USER IP ", []string{userIP})
	var request models.RegisterOwnerRequest

	// validating JSON body
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You didn't fill all the fields"})
		fmt.Println("ERROR", err)
		return
	}

	// checking if user already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", request.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find the user"})
		return
	}

	// hash password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Problem with registration"})
		return
	}

	// create new user
	user := models.User{
		FullName:     request.FullName,
		Email:        request.Email,
		Country:      request.Country,
		Address:      request.Address,
		PasswordHash: &hashedPassword,
		Role:         "owner",
		//ClientIP:     userIP, // wrap in slice
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		LastLogin:     nil,
		TermsAccepted: request.TermsAccepted,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already existing"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, "", user.Role, "auth", 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Activation link failed to generate"})
		return
	}

	_, err = EmailController.VerifyAccount(user.Email, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Activation link failed to send email"})
		return
	}

	// success
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
