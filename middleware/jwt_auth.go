package middleware

import (
	"net/http"
	"strings"

	"github.com/devfurkankizmaz/go-user-auth-service/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := utils.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := utils.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Not authorized"})
		c.Abort()
	}
}
