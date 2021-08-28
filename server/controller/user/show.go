package user

import (
	"github.com/gin-gonic/gin"
)

func ShowMassage(c *gin.Context) {
	massage := c.Param("massage")
	c.JSON(200, gin.H{
		"massage": massage,
	})
}
