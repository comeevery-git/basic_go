package types

/**
    어플리케이션 전체에서 공통적으로 사용되는 유틸
        - 문자열 처리, 날짜 형식 변환 등 범용 함수들
**/

import (
	"net/http"
	"encoding/json"
)

// HTTP 응답에 메시지를 작성하는 함수
func WriteResponse(w http.ResponseWriter, statusCode int, message string) {
    w.WriteHeader(statusCode)
    w.Write([]byte(message))
}

func HandleServiceFunc(w http.ResponseWriter, r *http.Request, serviceFunc func() (interface{}, error)) (interface{}, error) {
    data, err := serviceFunc()
    if err != nil {
        WriteResponse(w, http.StatusInternalServerError, err.Error())
        return nil, err
    }

    jsonData, err := json.Marshal(data)
    if err != nil {
        WriteResponse(w, http.StatusInternalServerError, err.Error())
        return nil, err
    }

    WriteResponse(w, http.StatusOK, string(jsonData))
    return data, nil
}