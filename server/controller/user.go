package controller

import (
	"fmt"
	"server/dbmod"
	"server/model"

	"github.com/gin-gonic/gin"
)

// func NewUser(c *gin.Context) {
// 	user := &model.User{}
// 	err := c.BindJSON(&user)
// 	if err != nil {
// 		fmt.Println(http.StatusBadRequest, "Request is failed: "+err.Error())
// 	}

// }

func InsertUser(c *gin.Context, registerUser *model.User, user_name string) {
	db := dbmod.SqlConnect()
	// insert
	result := db.Create(&registerUser)
	// user_back := []model.User{}
	// db.Where("name = ?", user_name).First(&user_back)
	err := result.Error
	if err != nil {
		fmt.Println("負け")
		c.JSON(400, err)
	} else {
		c.JSON(200, result)
	}
	// fmt.Println(reflect.TypeOf(err))

	// fmt.Println(err)
	defer db.Close()
}

func AddUser(c *gin.Context) {
	user_name := c.PostForm("username")
	pass := c.PostForm("pass")

	user := model.User{
		Name: user_name,
		Pass: pass,
	}

	InsertUser(c, &user, user_name)
}
