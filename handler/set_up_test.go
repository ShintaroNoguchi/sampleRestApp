package handler

import (
	"sampleRestApp/model"
)

type errorResponse struct {
	Error string
}

// usersのテスト用データ作成
func seedUsers() []model.User {
	return []model.User{
		{
			Name: "John Smith",
			Age:  22,
		},
		{
			Name: "Yamada Taro",
			Age:  33,
		},
	}
}
