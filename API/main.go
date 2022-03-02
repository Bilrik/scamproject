package main

import (
	"grpci"

	"github.com/gin-gonic/gin"
)

func setuproutes(app *gin.Engine) {
	app.GET("/api/call-protection/calls", grpci.calls)
	app.GET("/api/call-protection/batch", grpci.batch)
	app.POST("/api/v1/book", grpci.NewBook)
}

func main() {
	app := gin.Default()
	setuproutes(app)
}
