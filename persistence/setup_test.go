package persistence

import (
	"os"
	"sampleRestApp/db"
	"sampleRestApp/model"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

const location = "Asia/Tokyo"

// NewTestRepository test用DBを準備する
func NewTestRepository() (db.Repository, func(), error) {
	// test用の設定ファイルを読み込む
	os.Setenv("ENV", "persistence_test")

	// test用DBへ接続する
	repo, cleanup, err := db.NewRepository()
	if err != nil {
		return nil, nil, err
	}

	// cleanup db for test
	err = repo.GetConn().DropTableIfExists(&model.User{}).Error
	if err != nil {
		return nil, nil, err
	}

	// create tables for test
	err = repo.GetConn().AutoMigrate(
		model.User{},
	).Error
	if err != nil {
		return nil, nil, err
	}
	return repo, cleanup, nil
}

func seedUsers() ([]model.User, error) {
	users := []model.User{
		{
			Name: "John Smith",
			Age:  22,
		},
		{
			Name: "Yamada Taro",
			Age:  33,
		},
	}
	return users, nil
}

func insertUsers(r db.Repository) ([]model.User, error) {
	users, err := seedUsers()
	if err != nil {
		return nil, err
	}
	for _, v := range users {
		err = r.GetConn().Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	var insertedUsers []model.User
	err = r.GetConn().Find(&insertedUsers).Error
	if err != nil {
		return nil, err
	}
	return insertedUsers, nil
}

// now timezoneを指定した現在のtimeを返す
func now() time.Time {
	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Fatal(err)
	}
	return time.Now().In(loc)
}

// convertDBTime ミリ秒を切り捨てる
func convertDBTime(t time.Time) (*time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Fatal(err)
	}
	const layout = "2006-01-02 15:04:05 -0700 MST"
	t, parseErr := time.Parse(layout, t.Format(layout))
	if parseErr != nil {
		return nil, parseErr
	}
	t = t.In(loc)
	return &t, nil

}
