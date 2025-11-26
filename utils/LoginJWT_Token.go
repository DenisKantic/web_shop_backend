package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID   uint   `json:"user_id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Purpose  string `json:"purpose"` // "activation", "forgot_password", "auth"
	jwt.StandardClaims
}

func GenerateJWT(userID uint, fullname string, role string, purpose string, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration * time.Minute)

	claims := &Claims{
		FullName: fullname,
		UserID:   userID,
		Role:     role,
		Purpose:  purpose,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
