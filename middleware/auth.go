package middleware

import (
	"net/http"
	"os"
	"strings"

	s "github.com/firasam/julo/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(jwtService s.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var SECRET_KEY = []byte(os.Getenv("SECRET"))
		encodedToken := strings.Split(authHeader, " ")[1]
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(encodedToken, claims, func(token *jwt.Token) (interface{}, error) {
			return SECRET_KEY, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("xid", claims["xid"])
		c.Next()
	}
}
