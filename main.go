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

	var count int
	count = 1000

	err, message := getMessageFromJson(c)
	if err != nil {
		fmt.Println(err)
	}

	var total_elapsedTime_proto int
	var total_elapsedTime_json int

	for i := 0; i < count; i++ {
		startTime_proto := time.Now()
		data_proto, err_proto := getProto(message)
		if err_proto != nil {
			fmt.Println("proto err", err_proto)

		}
		key_proto := fmt.Sprintf("protoMessage%v", i)
		services.SetRedis(key_proto, data_proto)
		elapsedTime_proto := time.Since(startTime_proto).Microseconds()
		fmt.Println("time for proto is ", elapsedTime_proto)
		total_elapsedTime_proto += int(elapsedTime_proto)

		startTime_json := time.Now()
		data_json, err_json := getJson(message)
		if err_json != nil {
			fmt.Println("json err", err_json)

		}
		key_json := fmt.Sprintf("jsonMessage%v", i)
		services.SetRedis(key_json, data_json)
		elapsedTime_json := time.Since(startTime_json).Microseconds()
		fmt.Println("time for json is ", elapsedTime_json)
		total_elapsedTime_json += int(elapsedTime_json)

	}
	result := make(map[string]interface{})

	result["total_elapsedTime_json"] = total_elapsedTime_json
	result["total_elapsedTime_proto"] = total_elapsedTime_proto
	result["average_json"] = total_elapsedTime_json / count
	result["average_proto"] = total_elapsedTime_proto / count

	c.JSON(200, gin.H{"resilt": result})

}

func benchGetHandler(c *gin.Context) {

}
func protoSetHandler(c *gin.Context) {

	// Start measuring time
	startTime := time.Now()

	err, message := getMessageFromJson(c)
	if err != nil {
		fmt.Println("Message from json in proto set handler error ", err)
	}

	data, err := getProto(message)
	if err != nil {

		fmt.Println("error in getProto", err)
	}
	elapsedTime := time.Since(startTime)

	_, err = services.SetRedis("protoMessage", data)
	if err != nil {

		fmt.Println("Redis Err is ", err)
	}
	fmt.Println("Redis Err is", err)

	// Calculate the time taken

	// Create a JSON response
	response := map[string]interface{}{
		"message":    "Set Proto called",
		"time_taken": elapsedTime.Microseconds(), // Convert the time.Duration to a string
	}

	c.JSON(http.StatusOK, response)

}

func jsonSetHandler(c *gin.Context) {
	startTime := time.Now()

	err, message := getMessageFromJson(c)
	if err != nil {
		fmt.Println("Message from json in json set handler error ", err)
	}

	data, err := getJson(message)
	if err != nil {

		fmt.Println("error in geJson", err)
	}
	elapsedTime := time.Since(startTime)

	_, err = services.SetRedis("jsonMessage", data)
	if err != nil {

		fmt.Println("Redis Err is ", err)
	}

	response := map[string]interface{}{
		"message":    "Set json called",
		"time_taken": elapsedTime.Microseconds(), // Convert the time.Duration to a string
	}

	c.JSON(http.StatusOK, response)

}

func protoGetHandler(c *gin.Context) {

}

func jsonGetHandler(c *gin.Context) {

}

func getProto(message *message_proto.Message) (data []byte, err error) {

	// fmt.Println(" error in binding json inside proto ", err)

	data, err = proto.Marshal(message)
	if err != nil {
		panic(err)
	}
	// fmt.Println(data)
	return
}

func getJson(message *message_proto.Message) (data []byte, err error) {

	data, err = json.Marshal(message)
	if err != nil {
		panic(err)
	}
	// fmt.Println(data)
	return
}

func getMessageFromJson(c *gin.Context) (err error, message *message_proto.Message) {

	err = c.BindJSON(&message)
	return
}
