package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main(){
	r := gin.Default()
	r.GET("/", helloWorld)
	r.GET("/data", myData)
	r.Run(":3000")
}

func helloWorld(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"Welcome to Golang",
	})
}

func myData(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"id": 1,
		"name": "Renes Dwi Riwanto",
	})
}