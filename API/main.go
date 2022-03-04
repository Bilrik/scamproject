package main

import (
	"scamProtec/API/grpci"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/api/call-protection/batch/:Number", grpci.Batch)
	app.Run("localhost:50054")
}
