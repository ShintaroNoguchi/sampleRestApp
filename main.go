package main

import (
	"fmt"
	"sampleRestApp/db"
	"sampleRestApp/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router, cleanup, err := initialize()
	if err != nil {
		fmt.Printf("server start failed. %v", err)
	}
	defer cleanup()

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initialize() (*gin.Engine, func(), error) {
	repository, cleanup, err := db.NewRepository()
	if err != nil {
		return nil, nil, err
	}
	userHandler := handler.NewUserHandler(repository)

	// ルーティング
	r := gin.Default()

	r.GET("/users", userHandler.GetAllUser)
	//r.POST("/users", handler.PostUser())

	return r, cleanup, nil
}
