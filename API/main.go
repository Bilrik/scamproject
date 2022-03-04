package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.POST("/api/call-protection/batch", func(c *gin.Context) {
		ANumber := c.PostForm("ANumber")
		BNumber := c.PostForm("BNumber")

		c.JSON(200, gin.H{
			"status":  "posted",
			"ANumber": ANumber,
			"BNumber": BNumber,
		})
	})

	app.POST("/form_post", func(c *gin.Context) {

		test := c.PostForm("message")
		nick := c.PostForm("nick")
		fmt.Println(test)
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": test,
			"nick":    nick,
		})
	})

	app.Run("localhost:50054")
}
