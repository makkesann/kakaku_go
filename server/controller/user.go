package controller

import (
	"fmt"
	"net/http"
	"server/dbmod"
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

func InsertUser(registerUser *model.User) {
	db := dbmod.SqlConnect()
	// insert
	db.Create(&registerUser)
	defer db.Close()
}

func AddUser(c *gin.Context) {
	user_name := c.PostForm("username")
	pass := c.PostForm("pass")

	var user = model.User{
		Name: user_name,
		Pass: pass,
	}

	InsertUser(&user)
}

func ShowUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"massage": "ping",
	})
}
