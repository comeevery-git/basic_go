package network

import (	
	"net/http"
	"log"
	"example.com/m/config"
	"example.com/m/service"
	"example.com/m/types"
)

type Server struct {
	userService *service.UserService
	productService *service.ProductService
}

func NewServer(userService *service.UserService, productService *service.ProductService) *Server {
	return &Server{
		userService: userService,
		productService: productService,
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
		_, err := types.HandleServiceFunc(w, r, func() (interface{}, error) {
			return s.userService.GetAllUsers()
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	// 엔드포인트 처리 추가
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		_, err := types.HandleServiceFunc(w, r, func() (interface{}, error) {
			return s.productService.GetAllProducts()
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	// http.ListenAndServe 함수를 사용하여 HTTP 서버를 시작
	if err := http.ListenAndServe(config.ServerPort, nil); err != nil {
		log.Fatal(err)
	}
}

/*
	- Java 와 달리 Go 는 함수 오버로딩을 지원하지 않음
	- http.Request는 Go 의 인터페이스 타입 (Java 인터페이스와 비슷)
*/
func handleRequest(w http.ResponseWriter, r *http.Request) {
	message := "NOT SUPPORTED." // 변수 선언 및 초기화 TODO error message
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