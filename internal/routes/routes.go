package routes

import (
	"job/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter (router *gin.Engine){

	//API v1 group
	{	v1 := router.Group("/api/v1")
		v1.GET("/students", controllers.GetStudents)
		v1.GET("/students/:id", controllers.GetStudentsById)
		v1.POST("/students", controllers.AddStudent)
		v1.PUT("/students/:id",controllers.UpdateStudentData)
		v1.DELETE("/students:id", controllers.DeleteStudent)
	}

	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(http.StatusAccepted,gin.H{
			"message" : "pong",
		})
	})
}