package main

import (
	c "aws-wiper/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})

	r.GET("/jobs", c.ListJobs)
	r.GET("/jobs/:id", c.GetJob)
	r.POST("/jobs", c.CreateJob)
	r.PUT("/jobs/:id", c.UpdateJob)
	r.DELETE("/jobs/:id", c.DeleteJob)
}
