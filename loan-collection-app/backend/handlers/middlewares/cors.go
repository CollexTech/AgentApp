package middlewares

import (
	"backend/constants"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CustomCorsConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = append(config.AllowMethods, constants.AllowedCorsMethods...)
	config.AllowHeaders = append(config.AllowHeaders, constants.AllowedCorsHeaders...)
	config.AllowAllOrigins = true
	return cors.New(config)
}

func SetCorsHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
