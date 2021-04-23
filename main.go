package main

import (
	"ColdMountain/conf"
	_ "ColdMountain/conf"
	_ "ColdMountain/connection"
	"ColdMountain/graphql/adaptation"
	"ColdMountain/middlewares"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func getStructedMsgs(c *gin.Context) {
	// 升级 get 请求为 webSocket 协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	// Kafka Consumer
	coldQ := conf.GetGlobalConfig().ColdQ
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2
	consumer, err := sarama.NewConsumer([]string{coldQ.Address+":"+coldQ.Port}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition("000", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partitionConsumer.Close()

	for {
		needBreak := false
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
			//time.Sleep(time.Second*3)
			err = ws.WriteMessage(1, msg.Value)
			//fmt.Printf("%v\n", err)
			if err != nil {
				needBreak = true
			}
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		case <-time.After(30 * time.Second): // 超时断开 ws
			err = ws.WriteMessage(1, []byte("ping")) // ping 尝试
			if err != nil {
				needBreak = true
			}
		}
		if needBreak {
			break
		}
	}

	fmt.Printf("getStructedMsgs end\n")
}

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
		apiR.GET("/getStructedMsgs", getStructedMsgs)
	}
	_ = r.Run(":8899")
}
