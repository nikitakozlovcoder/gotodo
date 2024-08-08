package middlewares

import (
	"github.com/gin-gonic/gin"
	"gotodo/app/services"
	"log"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	jwtService services.IJwtService
}

func NewAuthMiddleware(jwtService services.IJwtService) *AuthMiddleware {
	return &AuthMiddleware{jwtService: jwtService}
}

func (m *AuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("AuthMiddleware")
		header := c.GetHeader("Authorization")

		if len(header) == 0 || !strings.Contains(header, "Bearer") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		header = strings.TrimPrefix(header, "Bearer ")
		userId, err := m.jwtService.ParseJwt(header)

		log.Printf("User authenticated: %v\n", userId)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userId", userId)

		c.Next()
	}
}
