package handler

import (
	"net/http"
	"sampleRestApp/db"
	"sampleRestApp/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// UserHandler ユーザ情報用のインターフェース
type UserHandler interface {
	GetAllUser(*gin.Context)
	CreateUser(*gin.Context)
}

type userPersistence struct {
	db *gorm.DB
}

// NewUserHandler 新しいUserHandlerを作成する
func NewUserHandler(r db.Repository) UserHandler {
	return &userPersistence{r.GetConn()}
}

// GetAllUser ユーザ情報を全件取得
func (up userPersistence) GetAllUser(c *gin.Context) {
	var users []model.User
	err := up.db.Find(&users).Error
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "internal_server_error")
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser 新しいユーザを作成
func (up userPersistence) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, "bad_request")
		return
	}

	err := up.db.Create(&user).Error
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "internal_server_error")
		return
	}

	c.JSON(http.StatusOK, user)
}
