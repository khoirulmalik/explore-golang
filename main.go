package main

import (
	"github.com/gin-gonic/gin"

	"job/internal/routes"
)


func main (){
	router := gin.Default()
	routes.SetupRouter(router)

	router.Run()
}