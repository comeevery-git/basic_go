package types

import (
	"net/http"
)

// HTTP 응답에 메시지를 작성하는 함수
func WriteResponse(w http.ResponseWriter, message string) {
	w.Write([]byte(message))
}

