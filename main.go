package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main(){
	r := gin.Default()
	r.GET("/", helloWorld)
	r.GET("/data", myData)
	r.GET("/data/:id", dataSpesific)
	r.GET("/query", dataQuery)
	r.Run(":3000")
}
func dataQuery(c *gin.Context){
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"message": "tes",
	})
}

func dataSpesific(c *gin.Context){
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
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