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

type userHandler struct {
	db *gorm.DB
}

// NewUserHandler 新しいUserHandlerを作成する
func NewUserHandler(r db.Repository) UserHandler {
	return &userHandler{r.GetConn()}
}

// GetAllUser ユーザ情報を全件取得
func (uh userHandler) GetAllUser(c *gin.Context) {
	var users []model.User
	err := uh.db.Find(&users).Error
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "internal_server_error")
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser 新しいユーザを作成
func (uh userHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, "bad_request")
		return
	}

	err := uh.db.Create(&user).Error
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "internal_server_error")
		return
	}

	c.JSON(http.StatusCreated, user)
}
