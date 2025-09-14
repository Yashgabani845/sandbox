package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	go worker()
	router := gin.Default()
	router.POST("/jobs", cretaeJob)
	router.GET("/job/:id", Getjob)

	router.Run(":8080")

}
