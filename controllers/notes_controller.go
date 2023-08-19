package controllers

import (
	"fmt"
	"github.com/behryuz/gin_notes/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NotesIndex(c *gin.Context) {
	notes := models.NotesAll()
	c.HTML(
		http.StatusOK,
		"notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

func NotesNew(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"notes/new.html",
		gin.H{},
	)
}

func NotesCreate(c *gin.Context) {
	name := c.PostForm("name")
	content := c.PostForm("content")
	models.NoteCreate(name, content)
	c.Redirect(http.StatusMovedPermanently, "notes")
}

func NotesShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	note := models.NoteFind(id)
	c.HTML(
		http.StatusOK,
		"notes/show.html",
		gin.H{
			"note": note,
		},
	)
}
