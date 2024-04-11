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

/**
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
