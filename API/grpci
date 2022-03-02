package grpci

import (
	"github.com/gin-gonic/gin"
)

func calls(c *gin.Context) {
	c.SEND()
}

func batch(c *gin.Context) {
	c.SEND("A single book")
}

func update(c *gin.Context) {
	c.SEND("update")
}
