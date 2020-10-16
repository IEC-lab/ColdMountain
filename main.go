package main

import (
	_ "ColdMountain/conf"
	_ "ColdMountain/connection"
	"ColdMountain/graphql/adaptation"
	"ColdMountain/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	apiR := r.Group("/api")
	{
		apiR.Use(middlewares.Allow())
		apiR.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		apiR.GET("/graphql/playground", adaptation.PlaygroundHandler())
		apiR.POST("/graphql/query", adaptation.GraphqlHandler())
	}
	_ = r.Run(":8899")
}
