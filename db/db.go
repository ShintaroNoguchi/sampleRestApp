package db

import (
	"sampleRestApp/model"

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

	// usersテーブルがなかった場合、マイグレーションとシーディングを実行
	if db.HasTable(&model.User{}) == false {
		db.AutoMigrate(&model.User{})

		var u1 = model.User{Name: "taro", Age: 18}
		db.Create(&u1)
		var u2 = model.User{Name: "jiro", Age: 22}
		db.Create(&u2)
	}

	return &repositoryStruct{db: db}, cleanup, nil
}

func (r *repositoryStruct) GetConn() *gorm.DB {
	return r.db
}
