package main

import (
	"github.com/behryuz/gin_notes/controllers"
	"github.com/behryuz/gin_notes/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/**")

	models.ConnectDatabase()
	models.DBMigrate()

	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new", controllers.NotesNew)
	r.POST("/notes", controllers.NotesCreate)
	r.GET("/notes/:id", controllers.NotesShow)

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Hello Gin",
		})
	})

	log.Println("Server Started!")

	r.Run()
}
