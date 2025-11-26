package models

import "time"

type User struct {
	ID              uint    `gorm:"primary_key"`
	FullName        string  `gorm:"size:50;not null"`
	Email           string  `gorm:"uniqueIndex;not null" binding:"required"`
	PasswordHash    *string // pointer: nil if using OAuth
	Role            string  `gorm:"role;not null"`
	CreatedAt       time.Time
	IsEmailVerified bool `gorm:"default:false"`
	UpdatedAt       time.Time
	PostCount       int        `gorm:"default:0"`
	LastLogin       *time.Time `gorm:"default:null"`
	TermsAccepted   string     `gorm:"not null"`
}

// ManualLoginUserRequest json validation object for login
type ManualLoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ChangePasswordRequest json validation object for changing password
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=8"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type DeleteUserAccountRequest struct {
	Password string `json:"password" binding:"required"`
}
