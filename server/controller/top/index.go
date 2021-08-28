package top

import (
	"github.com/gin-gonic/gin"
)

func IndexDisplayAction(c *gin.Context) {
	c.JSON(200, gin.H{
		"massage": "ping",
	})
}
