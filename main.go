package main

import (
	"embed"
	"net/http"
	"os"

	"aws-wiper/libs"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//go:embed public
var publicFs embed.FS

func main() {
	godotenv.Load()
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.ForceConsoleColor()

	r := gin.Default()
	r.Use(static.Serve("/", libs.EmbedFolder(publicFs, "public")))
	InitRoutes(r.Group("/api"))
	r.NoRoute(func(c *gin.Context) { c.FileFromFS("public/", http.FS(publicFs)) })
	r.Run(":" + os.Getenv("APP_PORT"))
}
