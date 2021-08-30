package controller

import (
	"fmt"
	"net/http"
	"server/model"

	"github.com/gin-gonic/gin"
)

func NewUser(c *gin.Context) {
	user := &model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

}

func ShowUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"massage": "ping",
	})
}
