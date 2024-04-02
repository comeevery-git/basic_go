package usecase

// (API Reqeust) Web(network.go) => Controller => Usecase => Repository

import (
	"example.com/m/domain/repository"
	"example.com/m/internal/adapter/presenter"
)

// java/spring 의 @Service 유사한 역할
// Go 에서는 클래스 대신 구조체를 사용하며, 메서드는 이 구조체에 연결된다.
type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

// As-Is (presenter 사용 전)
// func (s *UserUsecase) GetUser(id int) (*model.User, error) {
// 	user, err := s.repo.FindByID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if user == nil {
// 		return nil, errors.New("User not found")
// 	}
// 	return user, nil
// }

// To-Be (presenter 사용 후)
func (u *UserUsecase) GetUser(id int) (presenter.UserResponse, error) {
    user, err := u.repo.FindByID(id)
    if err != nil {
        return presenter.UserResponse{}, err
    }
    return presenter.NewUserResponse(user), nil
}

func (u *UserUsecase) GetAllUsers() ([]presenter.UserResponse, error) {
    users, err := u.repo.GetAllUsers()
    if err != nil {
        return nil, err
    }

    var responses []presenter.UserResponse
    for _, user := range users {
        responses = append(responses, presenter.NewUserResponse(user))
    }
    return responses, nil
}
