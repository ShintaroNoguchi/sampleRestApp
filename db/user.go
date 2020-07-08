package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Age  int    `form:"age" json:"age" xml:"age"  binding:"required"`
}

type UserRepository interface {
	GetAllUser(*gin.Context)
}

type userPersistence struct {
	db *gorm.DB
}

// NewUserPersistence 新しいUserRepositoryを作成する
func NewUserPersistence(r Repository) UserRepository {
	ur := &userPersistence{r.GetConn()}

	// usersテーブルがなかった場合、マイグレーションとシーディングを実行
	if ur.db.HasTable(&User{}) == false {
		ur.db.AutoMigrate(&User{})

		var u1 = User{Name: "taro", Age: 18}
		ur.db.Create(&u1)
		var u2 = User{Name: "jiro", Age: 22}
		ur.db.Create(&u2)
	}

	return ur
}

func (up userPersistence) GetAllUser(c *gin.Context) {
	var users []User
	err := up.db.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

// ----- 以下、工事中 -----

// PostUser userを保存
func PostUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}

		db, err := gorm.Open("mysql", "user:password@(mysql:3306)/sample?parseTime=true")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "DB connection error"})
			return
		}
		defer db.Close()

		db.Create(&user)

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": ""})
	}
}
