package persistence

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUser_Success(t *testing.T) {
	// テスト用DBへ接続
	testRepo, cleanup, err := NewTestRepository()
	if err != nil {
		log.Fatal("faild create test repository:", err)
	}
	defer cleanup()

	// テストデータをInsert
	seedUsers, err := insertUsers(testRepo)
	if err != nil {
		log.Fatal("faild insert data", err)
	}

	// テスト実行
	bp := NewUserPersistence(testRepo)
	users, getErr := bp.GetAllUser()
	if getErr != nil {
		log.Fatal(getErr)
	}

	// 結果の検証
	// Insetされているものが全て返ってきているか
	assert.ElementsMatch(t, seedUsers, users)
}
