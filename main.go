package main

import (
	"fmt"

	"github.com/c-m3-codin/red_sonto/message_proto"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

func main() {

	r := gin.Default()
	r.GET("/ping", pongHandler)

}

func pongHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func protoSetHandler(c *gin.Context) {

}

func jsonSetHandler(c *gin.Context) {

}

func protoGetHandler(c *gin.Context) {

}

func jsonGetHandler(c *gin.Context) {

}

func setProto(val string) (data []byte, err error) {
	// var text = []byte(val)
	message := &message_proto.Message{
		Text: val,
	}
	data, err = proto.Marshal(message)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	return
}
