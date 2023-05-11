package middlewares

import (
	"net/http"
	"strings"

	"go-practice/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Query("token")
		if tokenString == "" {
			bearerToken := c.Request.Header.Get("Authorization")
			if len(strings.Split(bearerToken, " ")) == 2 {
				tokenString = strings.Split(bearerToken, " ")[1]
			}

		}

		if tokenString == "" {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		err := token.ValidateToken(tokenString)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()

	}
}
