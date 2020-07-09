package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	myerrors "sampleRestApp/errors"
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

func TestGetAllUser_Failed_InternalServerError(t *testing.T) {
	// userPersistenceのMock作成
	userPersistence, cleanup := beforeEach(t)
	defer cleanup()

	seedError := myerrors.NewDB("test")
	// Mockのレスポンスを設定
	userPersistence.EXPECT().GetAllUser().Return(nil, seedError)

	// テストリクエスト実行
	router := initialize(userPersistence)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/users", nil)
	if err != nil {
		log.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// レスポンスのステータスが500か
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateUser_Success(t *testing.T) {
	// userPersistenceのMock作成
	userPersistence, cleanup := beforeEach(t)
	defer cleanup()

	// リクエストのデータを作成
	seedUser := model.User{
		Name: "Tokugawa Ieyasu",
		Age:  100,
	}
	// Mockの引数作成
	requestUser := model.User{
		Name: seedUser.Name,
		Age:  seedUser.Age,
	}
	// Mockのレスポンス作成
	expectedResponse := model.User{
		Name: seedUser.Name,
		Age:  seedUser.Age,
	}
	log.Debug(expectedResponse)
	// Mockのレスポンスを設定
	userPersistence.EXPECT().CreateUser(requestUser).Return(&expectedResponse, nil)

	// テストリクエスト実行
	router := initialize(userPersistence)
	w := httptest.NewRecorder()

	requestUserJSON, _ := json.Marshal(requestUser)
	bodyReader := strings.NewReader(string(requestUserJSON))
	req, err := http.NewRequest("POST", "/v1/users", bodyReader)
	if err != nil {
		log.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// JSONをUserにバインド
	var user model.User
	err = json.NewDecoder(w.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	// レスポンスのステータスコードは201か
	assert.Equal(t, http.StatusCreated, w.Code)
	// 返ってきたJSONが正しいか
	assert.Equal(t, user, seedUser)
}

func TestCreateUser_Failed_BadRequest(t *testing.T) {
	// userPersistenceのMock作成
	userPersistence, cleanup := beforeEach(t)
	defer cleanup()

	// テストリクエスト実行
	router := initialize(userPersistence)
	w := httptest.NewRecorder()
	bodyReader := strings.NewReader("hogehoge")
	req, err := http.NewRequest("POST", "/v1/users", bodyReader)
	if err != nil {
		log.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// レスポンスのステータスコードは400か
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateUser_Failed_InternalServerError(t *testing.T) {
	// userPersistenceのMock作成
	userPersistence, cleanup := beforeEach(t)
	defer cleanup()

	// リクエストのデータを作成
	seedUser := model.User{
		Name: "Tokugawa Ieyasu",
		Age:  100,
	}
	// Mockの引数作成
	requestUser := model.User{
		Name: seedUser.Name,
		Age:  seedUser.Age,
	}

	seedError := myerrors.NewDB("test")
	// Mockのレスポンスを設定
	userPersistence.EXPECT().CreateUser(requestUser).Return(nil, seedError)

	// テストリクエスト実行
	router := initialize(userPersistence)
	w := httptest.NewRecorder()

	requestUserJSON, _ := json.Marshal(requestUser)
	bodyReader := strings.NewReader(string(requestUserJSON))
	req, err := http.NewRequest("POST", "/v1/users", bodyReader)
	if err != nil {
		log.Fatal(err)
	}
	router.ServeHTTP(w, req)

	// レスポンスのステータスコードは500か
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
