package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/c-m3-codin/red_sonto/message_proto"
	"github.com/c-m3-codin/red_sonto/services"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

func main() {

	services.InitRedis()

	defer services.CloseRedis()

	r := gin.Default()
	r.GET("/ping", pongHandler)
	r.POST("/set_proto", protoSetHandler)
	r.POST("/set_json", jsonSetHandler)
	r.GET("/get_proto", protoGetHandler)
	r.GET("/get_json", jsonGetHandler)
	r.POST("/set_bench", benchSetHandler)
	r.GET("/get_bench", benchGetHandler)

	err := r.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}

}

func pongHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func benchSetHandler(c *gin.Context) {

}

func benchGetHandler(c *gin.Context) {

}
func protoSetHandler(c *gin.Context) {

	// Start measuring time
	startTime := time.Now()

	data, err := getProto(c)
	fmt.Println(data, err)

	_, err = services.SetRedis("protoMessage", data)
	fmt.Println("Redis Err is", err)

	// Calculate the time taken
	elapsedTime := time.Since(startTime)

	// Create a JSON response
	response := map[string]interface{}{
		"message":    "Set Proto called",
		"time_taken": elapsedTime.String(), // Convert the time.Duration to a string
	}

	c.JSON(http.StatusOK, response)

}

func jsonSetHandler(c *gin.Context) {
	startTime := time.Now()
	data, err := getJson(c)
	fmt.Println(data, err)

	_, err = services.SetRedis("jsonMessage", data)
	fmt.Println("Redis Err is ", err)

	elapsedTime := time.Since(startTime)
	response := map[string]interface{}{
		"message":    "Set Proto called",
		"time_taken": elapsedTime.String(), // Convert the time.Duration to a string
	}

	c.JSON(http.StatusOK, response)

}

func protoGetHandler(c *gin.Context) {

}

func jsonGetHandler(c *gin.Context) {

}

func getProto(c *gin.Context) (data []byte, err error) {

	var message *message_proto.Message

	c.BindJSON(&message)

	data, err = proto.Marshal(message)
	if err != nil {
		panic(err)
	}
	// fmt.Println(data)
	return
}

func getJson(c *gin.Context) (data []byte, err error) {

	var message *message_proto.Message

	c.BindJSON(&message)

	data, err = json.Marshal(message)
	if err != nil {
		panic(err)
	}
	// fmt.Println(data)
	return
}
