package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/joho/godotenv"
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
	// .envからDB接続情報を取得
	var err error
	env := os.Getenv("ENV")
	switch env {
	case "persistence_test":
		err = godotenv.Load("./../db/env/test.env")
	default:
		err = godotenv.Load("./db/env/mysql.env")
	}
	if err != nil {
		log.Fatal(err)
	}

	dbURL := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := gorm.Open("mysql", dbURL)
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
