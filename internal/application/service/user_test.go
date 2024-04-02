package service

import (
    "testing"
    "example.com/m/model"
    "example.com/m/repository"
)

type MockUserRepository struct {
    repository.UserRepository
}

func (m *MockUserRepository) FindByID(id int) (*model.User, error) {
    // 테스트를 위한 가짜 사용자 데이터 반환
    return &model.User{ID: id, Name: "Test User"}, nil
}

/*
	testing.T: Go 테스트 프레임워크에서 제공하는 테스트 핸들러
*/
func TestGetUser(t *testing.T) {
    // MockUserRepository와 UserService 생성
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo)

    // GetUser 메서드 호출
    user, err := service.GetUser(1)
    if err != nil {
        t.Fatal(err) // 메세지를 출력하고 테스트를 종료시킴
    }

    t.Log("GetUser 호출 성공")

    t.Run("CheckUserID", func(t *testing.T) { // 'CheckUserID'라는 하위 테스트 생성
        // 반환된 사용자 ID 검증
        if user.ID != 1 {
            t.Errorf("expected ID %v, got %v", 1, user.ID) // 메세지를 출력하지만 테스트를 종료시키지 않음
			// %v와 같은 형식 지정자를 사용하여 변수를 문자열에 포함할 수 있음
        } else {
            t.Log("User ID is correct") // -v 옵션 사용 시 로그 출력
        }
    })

    t.Run("CheckUserName", func(t *testing.T) {
        // 반환된 사용자 이름 검증
        if user.Name != "Test User" {
            t.Errorf("expected Name %v, got %v", "Test User", user.Name)
        } else {
            t.Log("User Name is correct")
        }
    })
}