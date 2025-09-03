package main

import (
	"github.com/gin-gonic/gin"

	"job/data"
)

func main (){
	router := gin.Default()

	router.GET("/students", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"data" : data.Students,
		})
	})

	router.POST("/students", func (c *gin.Context)  {
		
	})

	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"message" : "pong",
		})
	})

	router.Run()

}