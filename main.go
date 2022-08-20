package main

import (
	"learnGoAPI/handler"

	"github.com/gin-gonic/gin"
)
func main(){
	r := gin.Default()
	r.GET("/", handler.HelloWorld)
	r.GET("/data", handler.MyData)
	r.POST("/data", handler.DataInputHandler)
	r.GET("/data/:id", handler.DataSpesific)
	r.GET("/query", handler.DataQuery)
	r.Run(":3000")
}
