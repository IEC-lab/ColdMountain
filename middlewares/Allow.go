package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Allow() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Next()
	}
}