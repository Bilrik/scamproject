package grpci

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Calls(c *gin.Context) {
	fmt.Println()
}

func Batch(c *gin.Context) {
	fmt.Println("A single book")
}

func Update(c *gin.Context) {
	fmt.Println("update")
}
