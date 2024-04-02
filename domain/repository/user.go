package repository

// 인터페이스로 데이터베이스와 상호작용하는 메서드 정의

import "example.com/m/domain/model"

type UserRepository interface {
	FindByID(id int) (*model.User, error)
	Save(user *model.User) error
	Update(user *model.User) error
	Delete(id int) error
	GetAllUsers() ([]*model.User, error)
}