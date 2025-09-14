package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var Jobqueue = make(chan string, 100)

func cretaeJob(c *gin.Context) {
	var payload struct {
		Language string `json:"language"`
		Source   string `json:"source"`
		Stdin    string `json:"stdin"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	job := Job{
		Id:        uuid.New().String(),
		Language:  payload.Language,
		Source:    payload.Source,
		Stdin:     payload.Stdin,
		Status:    "queued",
		CreatedAt: time.Now(),
	}
	mu.Lock()
	jobs[job.Id] = job
	mu.Unlock()
	Jobqueue <- job.Id
	c.JSON(http.StatusCreated, job)
}

func Getjob(c *gin.Context) {
	id := c.Param("id")
	mu.Lock()
	job, exists := jobs[id]
	mu.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}
	c.JSON(http.StatusOK, job)
}
