package main

import (
	"./grpci"
	"github.com/gin-gonic/gin"
)

func setuproutes(app *gin.Engine) {
	app.GET("/api/call-protection/calls", grpci.Calls)
	app.GET("/api/call-protection/batch", grpci.Batch)
	app.POST("/api/v1/book", grpci.Update)
}

func main() {
	app := gin.Default()
	setuproutes(app)
}
