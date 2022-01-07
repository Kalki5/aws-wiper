package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GET /jobs
// List all the jobs
func ListJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "data"})
}

// GET /jobs/:id
// Get the status of a Job
func GetJob(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "data " + c.Param("id")})
}

// POST /jobs
// Create a new Job
func CreateJob(c *gin.Context) {
	var input struct {
		Id   string `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}
	id := uuid.New().String()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "data" + id})
}

// PUT /jobs/:id
// Update a Job's Status
func UpdateJob(c *gin.Context) {
	var input struct {
		Id     string `json:"id"`
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "data " + c.Param("id")})
}

// DELETE /jobs/:id
// Delete an existing Job
func DeleteJob(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "data " + c.Param("id")})
}
