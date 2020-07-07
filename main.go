package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		db, err := gorm.Open("mysql", "user:password@(mysql:3306)/sample?parseTime=true")
		if err != nil {
			panic("データベースへの接続に失敗しました")
		}
		defer db.Close()

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
