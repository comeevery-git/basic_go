package usecase

// (API Reqeust) Web(network.go) => Controller => Usecase => Repository

import (
	"context"
	"errors"

	pb "example.com/m/proto"

	"example.com/m/domain/repository"
	"example.com/m/internal/adapter/presenter"
)

// java/spring 의 @Service 유사한 역할
// Go 에서는 클래스 대신 구조체를 사용하며, 메서드는 이 구조체에 연결된다.
type UserUsecase struct {
	repo repository.UserRepository
	pb.UnimplementedUserServiceServer // gRPC SERVER - UserServiceServer 인터페이스를 임베드, 서비스가 UserServiceServer 인터페이스 모든 메서드를 구현하지 않아도 컴파일 에러가 발생하지 않음
}

func NewUserUsecase(r repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

/*
    gRPC SERVER
 */

func (u *UserUsecase) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := u.repo.FindByID(int(req.Id))
	if err != nil {
		return nil, errors.New("GetUser Error!")
	}

	if user == nil {
		return nil, errors.New("GetUser Null!")
	}

	return &pb.GetUserResponse{
		User: presenter.ConvertUserToResponse(user),
	}, nil
}

func (u *UserUsecase) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	users, err := u.repo.GetAllUsers()
	if err != nil {
		return nil, errors.New("GetAllUsers Error!")
	}

	if users == nil {
		return nil, errors.New("GetAllUsers Null!")
	}

	var responses []*pb.User
	for _, user := range users {
		responses = append(responses, presenter.ConvertUserToResponse(user))
	}

	return &pb.GetAllUsersResponse{
		Users: responses,
	}, nil
}

func (u *UserUsecase) CreateUesr(ctx context.Context, req *pb.CreateUesrRequest) {
	user := &pb.User{
		Id:           req.User.Id,
		UserName:     req.User.UserName,
        UserEmail:    req.User.UserEmail,
        Password:     req.User.Password,
        Memo:         req.User.Memo,
        UseYn:        req.User.UseYn,
	}

	u.repo.Save(presenter.ConvertUserToModel(user))
}

func (u *UserUsecase) CreateUesrTestData(ctx context.Context, req *pb.CreateUesrTestDataRequest) {
    var wg sync.WaitGroup // WaitGroup: 고루틴이 모두 끝날 때까지 대기할 수 있는 기능 제공
    for i := 0; i < int(req.NumUsers); i++ {
        wg.Add(1) // 고루틴이 시작될 때마다 카운트를 1씩 증가. WaitGroup의 Add(): 카운트를 증가시키는 함수
        go func(i int) { // 고루틴 시작
            defer wg.Done() // 고루틴이 끝나면 카운트를 1씩 감소. WaitGroup의 Done(): 카운트를 감소시키는 함수
            user := &pb.User{
                Id:       int32(i),
                UserName: fmt.Sprintf("user%d", i),
                UserEmail: fmt.Sprintf("user%d@test.com", i),
                Password:  "test",
                Memo:      "test",
                UseYn:     "Y",
            }
            u.repo.Save(presenter.ConvertUserToModel(user))
        }
    }

    
}

/**
	HTTP SERVER
func (u *UserUsecase) GetUser(id int) (presenter.UserResponse, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		// return presenter.UserResponse{}, err
		return presenter.UserResponse{}, errors.New("GetUser Error!")
	}
	return presenter.NewUserResponse(user), nil
}

func (u *UserUsecase) GetAllUsers() ([]presenter.UserResponse, error) {
	users, err := u.repo.GetAllUsers()
	if err != nil {
		// return nil, err
		return nil, errors.New("GetAllUsers Error!")
	}

	var responses []presenter.UserResponse
	for _, user := range users {
		responses = append(responses, presenter.NewUserResponse(user))
	}
	return responses, nil
}
*/
