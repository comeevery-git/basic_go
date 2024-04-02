package presenter

import (
	"encoding/json"
	"net/http"

	"example.com/m/domain/model"
)

// UserResponse는 클라이언트에게 반환될 사용자 정보 응답 구조체입니다.
type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// 필요한 다른 필드들...
}

// NewUserResponse는 도메인 모델을 받아서 응답용 구조체로 변환합니다.
func NewUserResponse(user *model.User) UserResponse {
	return UserResponse{
		ID:   user.ID,
		Name: user.UserName,
		// TODO 필요한 필드 변환 처리...
	}
}

// WriteJSONResponse는 HTTP 응답으로 JSON 형식의 사용자 정보를 반환합니다.
func WriteJSONResponse(w http.ResponseWriter, statusCode int, response UserResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}