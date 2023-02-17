package Controllers

import (
	"github.com/gin-gonic/gin"
)

func GetFeed(c *gin.Context) {

	c.JSON(200, gin.H{"123": "123"})
}
