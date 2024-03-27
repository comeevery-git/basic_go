package service

// Network, repository 의 다리 역할
// (API Reqeust) Network => Service => Repository

import (
	"errors"
	"example.com/m/model"
	"example.com/m/repository"
)

// java/spring 의 @Service 유사한 역할
// Go 에서는 클래스 대신 구조체를 사용하며, 메서드는 이 구조체에 연결된다.
type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUser(id int) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]*model.User, error) {
    users, err := s.repo.GetAllUsers()
    if err != nil {
        return nil, err
    }
    if users == nil {
        return nil, errors.New("No users found")
    }
    return users, nil
}

