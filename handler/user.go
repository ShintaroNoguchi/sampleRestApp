package handler

import (
	"net/http"
	"sampleRestApp/db"
	"sampleRestApp/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetAllUser(*gin.Context)
}

type userPersistence struct {
	db *gorm.DB
}

// NewUserPersistence 新しいUserRepositoryを作成する
func NewUserPersistence(r db.Repository) UserRepository {
	return &userPersistence{r.GetConn()}
}

func (up userPersistence) GetAllUser(c *gin.Context) {
	var users []model.User
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

		var user model.User
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
