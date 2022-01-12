package controllers

import (
	. "aws-wiper/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// List all the jobs
func ListJobs(c *gin.Context) {
	jobs := JobList.ListJobs()
	transformedJobs := make([]Map, 0, len(jobs))
	for _, j := range jobs {
		transformedJobs = append(transformedJobs, j.ToMapMini())
	}
	c.JSON(http.StatusOK, Map{"jobs": transformedJobs})
}

// Get the status of a Job
func GetJob(c *gin.Context) {
	job, err := JobList.GetJob(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, transformError(err))
		return
	}
	c.JSON(http.StatusOK, job.ToMap())
}

// Create a new Job
func CreateJob(c *gin.Context) {
	var jobParam JobParam
	if err := c.ShouldBindJSON(&jobParam); err != nil {
		c.JSON(http.StatusBadRequest, transformError(err))
		return
	}

	job, err := NewJob(jobParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, transformError(err))
		return
	}
	JobList.AddJob(job)

	c.JSON(http.StatusOK, job.ToMap())
}

// Update a Job's Status
func UpdateJob(c *gin.Context) {
	var input struct {
		Action string `json:"action" binding:"required,oneof=pause resume"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, transformError(err))
		return
	}

	job, err := JobList.UpdateJob(c.Param("id"), input.Action)
	if err != nil {
		c.JSON(http.StatusBadRequest, transformError(err))
		return
	}

	c.JSON(http.StatusOK, job.ToMap())
}

// Delete an existing Job
func DeleteJob(c *gin.Context) {
	job, err := JobList.DeleteJob(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, transformError(err))
		return
	}
	c.JSON(http.StatusOK, job.ToMap())
}
