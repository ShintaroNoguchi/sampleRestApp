package main

import (
	"sampleRestApp/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type user struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "user:password@(mysql:3306)/sample?parseTime=true")
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}
	defer db.Close()

	// usersテーブルがなかった場合、マイグレーションを実行
	if db.HasTable(&user{}) == false {
		db.AutoMigrate(&user{})
	}

	r := gin.Default()

	r.GET("/users", handler.PostUser())
	r.POST("/users", handler.GetUser())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
