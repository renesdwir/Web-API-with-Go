package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
func main(){
	r := gin.Default()
	r.GET("/", helloWorld)
	r.GET("/data", myData)
	r.POST("/data", dataInputHandler)
	r.GET("/data/:id", dataSpesific)
	r.GET("/query", dataQuery)
	r.Run(":3000")
}

type DataInput struct {
	Id int
	Name string
	GradeClass string `json:"grade_class"`
}

func dataInputHandler(c *gin.Context){
	var data DataInput

	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": data.Id,
		"name": data.Name,
		"grade_class": data.GradeClass,
	})
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