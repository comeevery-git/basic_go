package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/m/internal/application/usecase"
)

type UserController struct {
    userUsecase usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: u}
}

func (c *UserController) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	responses, err := c.userUsecase.GetAllUsers()
	if err != nil {
		// TODO 공통 오류 처리 필요
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 성공적으로 사용자 정보를 조회한 경우, 결과를 JSON 형식으로 응답합니다.
	// TODO 공통 응답 처리 필요
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responses); err != nil {
		// JSON 인코딩 오류 처리 TODO 공통 오류 처리 필요
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (c *UserController) HandleGetUser(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr) // strconv Go 표준 라이브러리를 사용하여 문자열을 int 타입으로 변환
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    response, err := c.userUsecase.GetUser(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}