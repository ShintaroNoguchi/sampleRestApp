package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_persistence "sampleRestApp/mock/persistence"
	"sampleRestApp/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func newUserHandlerForMock(up *mock_persistence.MockUserPersistence) UserHandler {
	return &userHandler{
		userPersistence: up,
	}
}

func initialize(up *mock_persistence.MockUserPersistence) *gin.Engine {
	userHandler := newUserHandlerForMock(up)

	// ルーティング
	r := gin.Default()

	r.GET("/v1/users", userHandler.GetAllUser)
	r.POST("/v1/users", userHandler.CreateUser)
	r.PUT("/v1/users/:id", userHandler.UpdateUser)
	r.DELETE("/v1/users/:id", userHandler.DeleteUser)

	return r
}

func beforeEach(t *testing.T) (*mock_persistence.MockUserPersistence, func()) {
	// Mock用のおまじない
	ctrl := gomock.NewController(t)

	// userPersistenceのMock作成
	return mock_persistence.NewMockUserPersistence(ctrl), ctrl.Finish
}

func TestGetAllUser_Success(t *testing.T) {
	// userPersistenceのMock作成
	userPersistence, cleanup := beforeEach(t)
	defer cleanup()

	seedUsers := seedUsers()
	// Mockのレスポンスを設定
	userPersistence.EXPECT().GetAllUser().Return(seedUsers, nil)

	// テストリクエスト実行
	router := initialize(userPersistence)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/users", nil)
	if err != nil {
		log.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// JSONをUserにバインド
	var users []model.User
	err = json.NewDecoder(w.Body).Decode(&users)
	if err != nil {
		log.Fatal(err)
	}

	// レスポンスのステータスが200か
	assert.Equal(t, 200, w.Code)
	// 返ってきたJSONが正しいか
	assert.ElementsMatch(t, users, seedUsers)
}
