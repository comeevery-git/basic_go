package web

import (
	// "log"
	"net/http"

	// "example.com/m/internal/adapter/controller"
	"example.com/m/internal/application/usecase"
	// "example.com/m/pkg/infrastructure/config"
	// "github.com/gorilla/mux"
)

/**
	gRPC SERVER
*/
type Server struct {
	// gRPC에서는 컨트롤러를 인스턴스화하는 대신 서비스 구현체를 인스턴스화하고 이를 gRPC 서버에 등록한다.
	// 클라이언트는 서비스 메서드를 호출하여 원격 프로시저 실행
	userUsecase *usecase.UserUsecase
}

func NewServer(userUsecase *usecase.UserUsecase) *Server {
	return &Server{
		userUsecase: userUsecase,
	}
}

/*
- Go 에서는 함수 선언 시 func 키워드 사용
- 접근제어자는 변수 대소문자로 구분할 뿐임 (대문자: public, 소문자: private)
*/
/*
	HTTP SERVER


type Server struct {
	userController *controller.UserController
}

func NewServer(userController *controller.UserController) *Server {
	return &Server{
		userController: userController,
	}
}

func StartServer(s *Server) {
	r := mux.NewRouter()

	// gorilla/mux 라이브러리 활용하여 라우팅 정의, Spring 의 @RequestMapping 와 유사한 역할
	r.HandleFunc("/", handleHome)
	r.HandleFunc("/users", s.userController.HandleGetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", s.userController.HandleGetUser).Methods("GET")

	// 추가되지 않은 모든 경로에 대해 "NOT SUPPORTED" 메시지 반환
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "NOT SUPPORTED.", http.StatusNotFound)
	})

	// http.ListenAndServe 함수를 사용하여 HTTP 서버를 시작
	log.Println("Server is starting on", config.ServerPort)
	if err := http.ListenAndServe(config.ServerPort, r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
*/

/*
- Java 와 달리 Go 는 함수 오버로딩을 지원하지 않음
- http.Request는 Go 의 인터페이스 타입 (Java 인터페이스와 비슷)
*/
func handleHome(w http.ResponseWriter, r *http.Request) {
	message := "HELLO, WORLD!"
	w.Write([]byte(message)) // HTTP 요청 처리
}

func (s *Server) handleRequestWithService(w http.ResponseWriter, r *http.Request, serviceFunc func() ([]byte, error)) {
	// serviceFunc를 사용하여 서비스 로직을 처리하고 결과를 가져옴
	data, err := serviceFunc()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// HTTP 응답으로 데이터를 전송
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
