package network

import (	
	"net/http"
	"encoding/json"
	"example.com/m/config"
	"example.com/m/service"
)

type Server struct {
	userService *service.UserService
}

func NewServer(userService *service.UserService) *Server {
	return &Server{
		userService: userService,
	}
}

/*
	- Go 에서는 함수 선언 시 func 키워드 사용
	- 접근제어자는 변수 대소문자로 구분할 뿐임 (대문자: public, 소문자: private)
*/
func StartServer(s *Server) {
	// HandleRequest 로 라우팅, Spring 의 @RequestMapping 와 유사한 역할
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        s.handleUsers(w, r)
    })
	// http.ListenAndServe 함수를 사용하여 HTTP 서버를 시작
	http.ListenAndServe(config.ServerPort, nil)
}

/*
	- Java 와 달리 Go 는 함수 오버로딩을 지원하지 않음
	- http.Request는 Go 의 인터페이스 타입 (Java 인터페이스와 비슷)
*/
func handleRequest(w http.ResponseWriter, r *http.Request) {
	message := "Hello, World!" // 변수 선언 및 초기화
	w.Write([]byte(message)) // HTTP 요청 처리
}
func (s *Server) handleUsers(w http.ResponseWriter, r *http.Request) {
	// UserService를 사용하여 모든 사용자를 가져옴
	users, err := s.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 사용자 데이터를 JSON으로 변환
	userData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// HTTP 응답으로 사용자 데이터를 전송
	w.Header().Set("Content-Type", "application/json")
	w.Write(userData)
}