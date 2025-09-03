package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"job/data"
)

type CreateStudentInput struct {
    ID      int    `json:"id" binding:"required"`
    Nama    string `json:"nama" binding:"required"`
    NIM     string `json:"nim" binding:"required"`
    Jurusan string `json:"jurusan" binding:"required"`
}

type UpdateStudntInput struct {
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

	router.GET("/students/:id", func (c *gin.Context){
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

	  if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid Student ID",
        })
        return 
    	}

		c.JSON(http.StatusOK, gin.H{
			"data": data.Students[id],
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

router.PUT("/students/:id", func(c *gin.Context) {
    var input UpdateStudntInput
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr) 
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid Student ID",
        })
        return 
    }
    
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    for i := 0; i < len(data.Students); i++ {
        if data.Students[i].ID == id {
            data.Students[i] = data.Student{
                ID:      id, 
                Nama:    input.Nama,
                NIM:     input.NIM,
                Jurusan: input.Jurusan,
            }
            c.JSON(http.StatusOK, gin.H{
                "data": data.Students[i],
            })
            return
        }
    }


    c.JSON(http.StatusNotFound, gin.H{
        "error": "Student not found",
    })
})


	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"message" : "pong",
		})
	})

	router.Run()

}