package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetFeed(c *gin.Context) {
	logrus.Info("Get Feed!!!!!!!!!!!!!!!!!!!")

	c.JSON(200, gin.H{"123": "123"})
}
