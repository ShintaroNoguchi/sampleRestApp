package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Repository DB接続情報の構造体
type Repository interface {
	GetConn() *gorm.DB
}

type repositoryStruct struct {
	db *gorm.DB
}

// NewRepository DBに接続
func NewRepository() (Repository, func(), error) {
	db, err := gorm.Open("mysql", "user:password@(mysql:3306)/sample?parseTime=true")
	if err != nil {
		return nil, nil, err
	}

	// 接続切断用の関数
	cleanup := func() {
		db.Close()
	}

	return &repositoryStruct{db: db}, cleanup, nil
}

func (r *repositoryStruct) GetConn() *gorm.DB {
	return r.db
}
