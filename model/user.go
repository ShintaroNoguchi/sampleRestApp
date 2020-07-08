package model

import (
	"github.com/jinzhu/gorm"
)

// User ユーザ情報の構造体
type User struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}
