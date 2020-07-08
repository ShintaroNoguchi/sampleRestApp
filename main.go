package main

import (
	"fmt"
	"sampleRestApp/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type user struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	cleanup, err := initialize()
	if err != nil {
		fmt.Printf("server start failed. %v", err)
	}
	defer cleanup()
}

func initialize() (func(), error) {
	repository, cleanup, err := db.NewRepository()
	if err != nil {
		return nil, err
	}
	userRepository := db.NewUserPersistence(repository)

	// ルーティング
	r := gin.Default()

	r.GET("/users", userRepository.GetAllUser)
	//r.POST("/users", handler.PostUser())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	return cleanup, nil
}
