package handler

import (
	"fmt"
	"net/http"

	"learnGoAPI/dataInput"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func DataInputHandler(c *gin.Context){
	var data dataInput.DataInput

	err := c.ShouldBindJSON(&data)
	fmt.Println(err)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": data.Id,
		"name": data.Name,
		"grade_class": data.GradeClass,
	})
}



func DataQuery(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"name":    name,
		"message": "tes",
	})
}

func DataSpesific(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Golang",
	})
}

func MyData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   1,
		"name": "Renes Dwi Riwanto",
	})
}