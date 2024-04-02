package usecase

import (
	"testing"

	"example.com/m/domain/model"
	"example.com/m/domain/repository"
	"github.com/stretchr/testify/assert"
)

// UserRepository 인터페이스를 구현하는 Mock 객체 정의
type MockUserRepository struct {
	repository.UserRepository // UserRepository 인터페이스를 내장하여 인터페이스를 만족시킴
	FindByIDFunc func(int) (*model.User, error) // 모의 함수 작성
	GetAllUsersFunc func() ([]*model.User, error)
}

// FindByID 메서드 구현. FindByIDFunc 모의 함수 호출
func (mock *MockUserRepository) FindByID(id int) (*model.User, error) {
	return mock.FindByIDFunc(id)
}

// GetAllUsers 메서드 구현. GetAllUsersFunc 모의 함수 호출
func (mock *MockUserRepository) GetAllUsers() ([]*model.User, error) {
	return mock.GetAllUsersFunc()
}

// testing.T: Go 테스트 프레임워크에서 제공하는 테스트 핸들러
func TestUserUsecase_GetUser(t *testing.T) {
	// MockUserRepository 초기화 및 FindByID 메서드가 호출될 때의 동작 정의
	mockRepo := &MockUserRepository{
		FindByIDFunc: func(id int) (*model.User, error) {
			return &model.User{ID: id, UserName: "Lydia"}, nil
		},
	}
	userUsecase := NewUserUsecase(mockRepo) // 모의 리포지토리를 사용하여 UserUsecase 인스턴스를 생성

	response, err := userUsecase.GetUser(1)

    // github.com/stretchr/testify 라이브러리 활용하여 assert 패키지 사용
	assert.NoError(t, err) // 에러가 발생하지 않았는지 확인
	assert.Equal(t, "Lydia", response.Name)
}

func TestUserUsecase_GetAllUsers(t *testing.T) {
	mockRepo := &MockUserRepository{
		GetAllUsersFunc: func() ([]*model.User, error) {
			return []*model.User{{ID: 1, UserName: "Lydia"}}, nil // 한 명의 사용자만 포함하는 목록 반환
		},
	}
	userUsecase := NewUserUsecase(mockRepo)

	responses, err := userUsecase.GetAllUsers()

	assert.NoError(t, err)
	assert.Len(t, responses, 1)
	assert.Equal(t, "Lydia", responses[0].Name)
}