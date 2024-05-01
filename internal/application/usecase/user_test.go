package usecase

import (
	"context"
	"sync"
	"testing"

	"example.com/m/domain/model"
	"example.com/m/domain/repository"
	pb "example.com/m/proto"
	"github.com/stretchr/testify/assert"
)

// UserRepository 인터페이스를 구현하는 Mock 객체 정의
type MockUserRepository struct {
	repository.UserRepository                                // UserRepository 인터페이스를 내장하여 인터페이스를 만족시킴
	FindByIDFunc              func(int) (*model.User, error) // 모의 함수 작성
	GetAllUsersFunc           func() ([]*model.User, error)
	SaveFunc                  func(user *model.User) error
	calls                     int // Save 메소드 호출 횟수를 저장
	wg                        *sync.WaitGroup
}

// FindByID 메서드 구현. FindByIDFunc 모의 함수 호출
func (mock *MockUserRepository) FindByID(id int) (*model.User, error) {
	return mock.FindByIDFunc(id)
}

// GetAllUsers 메서드 구현. GetAllUsersFunc 모의 함수 호출
func (mock *MockUserRepository) GetAllUsers() ([]*model.User, error) {
	return mock.GetAllUsersFunc()
}
func (m *MockUserRepository) Save(user *model.User) error {
	if m.SaveFunc != nil {
		err := m.SaveFunc(user)
		m.calls++   // Save 메소드가 호출될 때마다 호출 횟수를 증가
		m.wg.Done() // Save 메소드가 호출될 때마다 Done 메소드를 호출
		return err
	}
	return nil
}
func (m *MockUserRepository) SaveCalls() int {
	return m.calls // Save 메소드가 호출된 횟수를 반환
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

	response, err := userUsecase.GetUser(context.Background(), &pb.GetUserRequest{Id: 1})

	// github.com/stretchr/testify 라이브러리 활용하여 assert 패키지 사용
	assert.NoError(t, err) // 에러가 발생하지 않았는지 확인
	assert.Equal(t, "Lydia", response.User.UserName)
}

func TestUserUsecase_GetAllUsers(t *testing.T) {
	mockRepo := &MockUserRepository{
		GetAllUsersFunc: func() ([]*model.User, error) {
			return []*model.User{{ID: 1, UserName: "Lydia"}}, nil // 한 명의 사용자만 포함하는 목록 반환
		},
	}
	userUsecase := NewUserUsecase(mockRepo)

	responses, err := userUsecase.GetAllUsers(context.Background(), &pb.GetAllUsersRequest{})

	assert.NoError(t, err)
	assert.Len(t, responses.Users, 1)
	assert.Equal(t, "Lydia", responses.Users[0].UserName)
}

func TestUserUsecase_CreateUser(t *testing.T) {
	mockRepo := &MockUserRepository{}
	userUsecase := NewUserUsecase(mockRepo)

	user := &pb.CreateUserRequest{
		User: &pb.User{
			Id:        1,
			UserName:  "Lydia",
			UserEmail: "test1234@test.com",
			Password:  "1234",
		},
	}
	userUsecase.CreateUser(context.Background(), user)

	assert.Equal(t, 1, user.User.Id)
	assert.Equal(t, "Lydia", user.User.UserName)
}

func TestUserUsecase_CreateUserTestData(t *testing.T) {
	var wg sync.WaitGroup
	mockRepo := &MockUserRepository{
		wg: &wg,
		SaveFunc: func(user *model.User) error {
			return nil
		},
	}
	userUsecase := NewUserUsecase(mockRepo)

	req := &pb.CreateUserTestDataRequest{
		NumUsers: 10,
	}

	wg.Add(int(req.NumUsers)) // 고루틴의 수만큼 카운트를 증가
	userUsecase.CreateUserTestData(context.Background(), req)
	wg.Wait() // 모든 고루틴이 완료될 때까지 대기

	// Save 메소드의 호출 횟수를 확인
	assert.Equal(t, int(req.NumUsers), mockRepo.SaveCalls())
}
