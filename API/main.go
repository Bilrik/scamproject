package main

import (
	"github.com/Bilrik/scamproject/grpci"

	"github.com/gin-gonic/gin"
)

func setuproutes(app *gin.Engine) {
	app.GET("/api/v1/book", grpci.GetBooks)
	app.GET("/api/v1/book/:id", grpci.GetBook)
	app.POST("/api/v1/book", grpci.NewBook)
}

func main() {
	app := gin.Default()
	setuproutes(app)
}
