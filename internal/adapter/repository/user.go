package repository

import (
	"github.com/jinzhu/gorm"

	"example.com/m/domain/model"
	"example.com/m/domain/repository"
)

/**
 *  gorm.DB: Gorm 라이브러리에서 제공하는 데이터베이스 연결 및 작업을 관리하는 주요 구조체
 */
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindByID(id int) (*model.User, error) {
	user := &model.User{}
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Save(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id int) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepository) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
