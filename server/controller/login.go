package controller

import (
	"fmt"
	"server/dbmod"
	"server/model"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	// username := c.PostForm("username")
	// pass := c.PostForm("pass")
	// fmt.Println(username)
	user := []model.Drink{}
	db := dbmod.SqlConnect()
	err := db.Where("ID = ?", "5").First(&user).Error
	fmt.Print(err)
	if err != nil {
		c.Status(400)
	}
	defer db.Close()
	// c.Status(400)

}
