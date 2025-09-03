package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"job/data"
)

type CreateStudentInput struct {
    ID      int    `json:"id" binding:"required"`
    Nama    string `json:"nama" binding:"required"`
    NIM     string `json:"nim" binding:"required"`
    Jurusan string `json:"jurusan" binding:"required"`
}

func main (){
	router := gin.Default()

	router.GET("/students", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"data" : data.Students,
		})
	})

	router.POST("/students", func (c *gin.Context)  {
		var input CreateStudentInput
		var newStudent data.Student
		if err := c.ShouldBindJSON(&input); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})

			return
		}

		input.ID = len(data.Students) + 1
		newStudent = data.Student{
			ID: input.ID,
			Nama: input.Nama,
			NIM: input.NIM,
			Jurusan: input.Jurusan,
		}

		data.Students = append(data.Students, newStudent)
		c.JSON(http.StatusCreated, gin.H{
			"data" : input,
		})
		
	})

	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"message" : "pong",
		})
	})

	router.Run()

}