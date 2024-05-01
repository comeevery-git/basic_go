package presenter

import (
	"encoding/json"
	"net/http"

	"example.com/m/domain/model"
	pb "example.com/m/proto"
)

// UserResponse는 클라이언트에게 반환될 사용자 정보 응답 구조체입니다.
type UserResponse struct {
	ID        int     `json:"id"`
	UserName  string  `json:"user_name"`
	UserEmail string  `json:"user_email"`
	Password  string  `json:"password"`
	Memo      *string `json:"memo,omitempty"`
	Status    string  `json:"status"`
}

/*
gRPC SERVER
*/
func ConvertUserToResponse(user *model.User) *pb.User {
	memo := ""
	if user.Memo != nil {
		memo = *user.Memo
	}

	if user == nil {
		return nil
	}

	return &pb.User{
		Id:        int32(user.ID),
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
		Password:  user.Password,
		Memo:      memo,
		Status:    user.Status,
	}
}

func ConvertUserToModel(user *pb.User) *model.User {
	return &model.User{
		ID:        int(user.Id),
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
		Password:  user.Password,
		Memo:      &user.Memo,
		Status:    user.Status,
	}
}

/**
	HTTP SERVER
// NewUserResponse는 도메인 모델을 받아서 응답용 구조체로 변환합니다.
func NewUserResponse(user *model.User) UserResponse {
	return UserResponse{
		ID:   user.ID,
		Name: user.UserName,
		// TODO 필요한 필드 변환 처리...
	}
}
*/

// WriteJSONResponse는 HTTP 응답으로 JSON 형식의 사용자 정보를 반환합니다.
func WriteJSONResponse(w http.ResponseWriter, statusCode int, response UserResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
