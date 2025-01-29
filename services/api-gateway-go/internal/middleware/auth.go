package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies the presence and validity of the Authorization header
func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	// Example: Validate the token (this should be replaced with real validation)
	if !strings.HasPrefix(authHeader, "Bearer ") || len(authHeader) < 8 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
		c.Abort()
		return
	}

	// TODO: Add token validation logic (e.g., JWT verification)

	c.Next()
}
