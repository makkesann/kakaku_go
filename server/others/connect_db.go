package others

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func sqlConnect() (database *gorm.DB) {
	HOST := "postgres"
	PORT := "5432"
	USER := "root"
	DBNAME := "test2"
	PASSWORD := "root"
	SSLMODE := "disable"
	CONNECT := "postgres://" + USER + ":" + PASSWORD + "@" + HOST + ":" + PORT + "/" + DBNAME + "?sslmode=" + SSLMODE
	db, err := gorm.Open(HOST, CONNECT)
	count := 0
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 50 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open(HOST, CONNECT)
		}
	}
	fmt.Println("DB接続成功")
	return db
}
