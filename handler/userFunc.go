package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type user struct {
	gorm.Model
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Age  int    `form:"age" json:"age" xml:"age"  binding:"required"`
}

// GetUser userを取得
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := gorm.Open("mysql", "user:password@(mysql:3306)/sample?parseTime=true")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "DB connection error"})
		}
		defer db.Close()

		users := []user{}
		db.Find(&users)

		c.JSON(http.StatusOK, users)
	}
}

// PostUser userを保存
func PostUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user user
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}

		db, err := gorm.Open("mysql", "user:password@(mysql:3306)/sample?parseTime=true")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "DB connection error"})
		}
		defer db.Close()

		db.Create(&user)

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": ""})
	}
}
