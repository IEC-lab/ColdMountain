package main

import (
	"ColdMountain/middlewares"
	"ColdMountain/streams"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	apiR := r.Group("/api")
	{
		apiR.Use(middlewares.Allow())
		apiR.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		apiR.GET("/streams", GetStreams)
	}
	_ = r.Run(":8888")
}

func GetStreams(context *gin.Context){
	context.JSON(200, gin.H{
		"streams": streams.DiscoverStreams(),
	})
}