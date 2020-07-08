package handler

import (
	"net/http"
	"sampleRestApp/db"
	"sampleRestApp/model"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// UserHandler ユーザ情報用のインターフェース
type UserHandler interface {
	GetAllUser(*gin.Context)
	CreateUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
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
		c.JSON(http.StatusBadRequest, err.Error())
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

// UpdateUser 新しいユーザ情報を更新
func (uh userHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var target model.User
	err = uh.db.First(&target, id).Error
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, "not_found")
		return
	} else if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "internal_server_error")
		return
	}

	target.Name = user.Name
	target.Age = user.Age
	err = uh.db.Save(&target).Error
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "internal_server_error")
		return
	}

	c.JSON(http.StatusOK, target)
}

// DeleteUser ユーザ情報を削除
func (uh userHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var user model.User
	err = uh.db.First(&user, id).Error
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, "not_found")
		return
	} else if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "internal_server_error")
		return
	}

	err = uh.db.Delete(&user).Error
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "internal_server_error")
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
