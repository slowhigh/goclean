package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsHandler() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowCredentials: true,
			AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
			AllowOrigins:     []string{"*"},
		})
}
