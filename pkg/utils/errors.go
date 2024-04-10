package utils

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func WrapWithMessage(err error, msg string) error {
	log.Errorf("WrapWithMessage: %+v\n", err)
	return wrap(err, msg)
}

func Wrap(err error) error {
	log.Errorf("Wrap: %+v\n", err)
	return wrap(err, "")
}

func wrap(err error, msg string) error {
	wrapFormat := "%s\n%w" // "{file:line} [func name] msg \n error"
	pc, file, line, ok := runtime.Caller(1)

	if !ok {
		return fmt.Errorf(wrapFormat, msg, err)
	}

	// {file:line} [funcName] msg
	stack := fmt.Sprintf("%s %s", getFuncInfo(pc, file, line), msg)
	return fmt.Errorf(wrapFormat, stack, err)
}

func getFuncInfo(pc uintptr, file string, line int) string {
	f := runtime.FuncForPC(pc) // pc 값을 사용하여 현재 실행 중인 함수의 정보를 얻음
	funcName := "unknown"      // 함수 이름을 저장할 변수, 기본값은 "unknown"
	if f != nil {
		funcName = f.Name() // f가 nil이 아니면, 함수의 이름을 가져옴
	}

	funcInfoFormat := "{%s:%d} [%s]"
	// 파일 이름, 라인 번호, 함수 이름을 포맷팅하여 문자열로 반환
	return fmt.Sprintf(funcInfoFormat, file, line, funcName)
}
