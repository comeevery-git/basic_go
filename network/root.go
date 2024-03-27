package network

import (
	"net/http"
	"example.com/m/config"
)

/*
	- Java 와 달리 Go 는 함수 오버로딩을 지원하지 않음
	- http.Request는 Go 의 인터페이스 타입 (Java 인터페이스와 비슷)
*/
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	message := "Hello, World!" // 변수 선언 및 초기화

	// HTTP 요청 처리
	w.Write([]byte(message))
}

/*
	- Go 에서는 함수 선언 시 func 키워드 사용
	- 접근제어자는 변수 대소문자로 구분할 뿐임 (대문자: public, 소문자: private)
*/
func StartServer() {
	// HandleRequest 로 라우팅, Spring 의 @RequestMapping 와 유사한 역할
	http.HandleFunc("/", HandleRequest)

	// http.ListenAndServe 함수를 사용하여 HTTP 서버를 시작
	http.ListenAndServe(config.ServerPort, nil)
}
