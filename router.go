package main

import (
	"aws-wiper/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})

	r.GET("/jobs", controllers.ListJobs)
	r.GET("/jobs/:id", controllers.GetJob)
	r.POST("/jobs", controllers.CreateJob)
	r.PUT("/jobs/:id", controllers.UpdateJob)
	r.DELETE("/jobs/:id", controllers.DeleteJob)
}
