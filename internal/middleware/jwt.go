package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jeevangb/project-portal-gateway/internal/auth"
)

type GraphQLRequest struct {
	Query string `json:"query"`
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow GraphQL Playground and Introspection Queries (GET requests)
		if c.Request.Method == http.MethodGet {
			c.Next()
			return
		}

		// Read request body safely
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			c.Abort()
			return
		}

		// Restore request body for future use
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// Parse GraphQL request
		var requestBody GraphQLRequest
		if err := json.Unmarshal(bodyBytes, &requestBody); err == nil {
			queryLower := strings.ToLower(requestBody.Query)

			// Allow unauthenticated access for signup and login
			if strings.Contains(queryLower, "mutation") &&
				(strings.Contains(queryLower, "signup") || strings.Contains(queryLower, "login")) {
				c.Next()
				return
			}
		}

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
		c.Next()
	}
}

func extractTokenFromHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
		return parts[1]
	}
	return ""
}
