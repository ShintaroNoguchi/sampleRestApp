package handler

import (
	"net/http"
	"sampleRestApp/model"
	"sampleRestApp/persistence"
	"strconv"

	"github.com/gin-gonic/gin"
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
	userPersistence persistence.UserPersistence
}

// NewUserHandler 新しいUserHandlerを作成する
func NewUserHandler(up persistence.UserPersistence) UserHandler {
	return &userHandler{
		userPersistence: up,
	}
}

// GetAllUser ユーザ情報を全件取得
func (uh userHandler) GetAllUser(c *gin.Context) {
	users, err := uh.userPersistence.GetAllUser()
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

	target, err := uh.userPersistence.CreateUser(user)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "internal_server_error")
		return
	}

	c.JSON(http.StatusCreated, *target)
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

	target, err := uh.userPersistence.UpdateUser(id, user)
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, "not_found")
		return
	} else if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "internal_server_error")
		return
	}

	c.JSON(http.StatusOK, *target)
}

// DeleteUser ユーザ情報を削除
func (uh userHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = uh.userPersistence.DeleteUser(id)
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, "not_found")
		return
	} else if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "internal_server_error")
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
