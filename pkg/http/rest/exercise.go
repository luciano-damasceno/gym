package rest

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World!"})
}

func os(c *gin.Context) {
	c.String(200, runtime.GOOS)
}
