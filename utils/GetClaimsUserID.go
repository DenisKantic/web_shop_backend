package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetClaims(c *gin.Context) (*Claims, bool) {

	claimsData, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return nil, false
	}

	claims, ok := claimsData.(*Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return nil, false
	}

	return claims, true
}
