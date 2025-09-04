package controllers

import (
	"net/http"
	"strconv"

	"job/data"

	"github.com/gin-gonic/gin"
)

type CreateStudentInput struct {
    ID      int    `json:"id" binding:"required"`
    Nama    string `json:"nama" binding:"required"`
    NIM     string `json:"nim" binding:"required"`
    Jurusan string `json:"jurusan" binding:"required"`
}

type UpdateStudentInput struct {
	Nama    string `json:"nama" binding:"required"`
    NIM     string `json:"nim" binding:"required"`
    Jurusan string `json:"jurusan" binding:"required"`
}

// GET - get all students controlleer
func GetStudents (c *gin.Context) {
	c.JSON(200,gin.H{
		"data" : data.Students,
	})
}

// GET - get a student by ID
func GetStudentsById (c *gin.Context){
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
}
// Post - add a student controller
func AddStudent (c *gin.Context){
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
		
}


// PUT - update data a student
func UpdateStudentData (c *gin.Context){
	var input UpdateStudentInput
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
}

// DELETE - delete student data 
func DeleteStudent (c *gin.Context){
	idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Student ID",
			})
			return
		}

		for i, s := range data.Students {
			if s.ID == id {
				data.Students = append(data.Students[:i], data.Students[i+1:]...)

				c.JSON(http.StatusOK, gin.H{
					"message": "student deleted successfully",
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})
}



