package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jeevangb/project-portal-gateway/internal/auth"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenstring := extractTokenFromHeader(c)
		if tokenstring == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}
		claims, err := auth.ParseToken(tokenstring)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		ctx := context.WithValue(c.Request.Context(), "claimsKey", claims)
		c.Request = c.Request.WithContext(ctx)
	}
}

func extractTokenFromHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}
