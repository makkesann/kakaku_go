package controller

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name     string    `json:"name"`
	Age      int       `json:age`
	Birthday time.Time `json:birthday`
}

func IndexDisplayAction(c *gin.Context) {
	c.JSON(200, gin.H{
		"massage": "ping",
	})
}
