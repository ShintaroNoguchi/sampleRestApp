package persistence

import (
	"sampleRestApp/db"
	"sampleRestApp/model"

	"github.com/jinzhu/gorm"
)

// UserPersistence ユーザ情報用のインターフェース
type UserPersistence interface {
	GetAllUser() ([]model.User, error)
	CreateUser(model.User) (*model.User, error)
	UpdateUser(uint64, model.User) (*model.User, error)
	DeleteUser(uint64) error
}

type userPersistence struct {
	db *gorm.DB
}

// NewUserPersistence 新しいUserPersistenceを作成する
func NewUserPersistence(r db.Repository) UserPersistence {
	return &userPersistence{r.GetConn()}
}

// GetAllUser ユーザ情報を全件取得
func (up userPersistence) GetAllUser() ([]model.User, error) {
	var users []model.User
	err := up.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser 新しいユーザを作成
func (up userPersistence) CreateUser(user model.User) (*model.User, error) {
	err := up.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser 新しいユーザ情報を更新
func (up userPersistence) UpdateUser(id uint64, user model.User) (*model.User, error) {
	var target model.User
	err := up.db.First(&target, id).Error
	if err != nil {
		return nil, err
	}

	target.Name = user.Name
	target.Age = user.Age
	err = up.db.Save(&target).Error
	if err != nil {
		return nil, err
	}

	return &target, nil
}

// DeleteUser ユーザ情報を削除
func (up userPersistence) DeleteUser(id uint64) error {
	var user model.User
	err := up.db.First(&user, id).Error
	if err != nil {
		return err
	}

	err = up.db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
