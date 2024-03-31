package types

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