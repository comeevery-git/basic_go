package experiment

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
    c := make(chan int)
    errC := make(chan error) // 오류를 수신할 채널 추가

    go func() {
        err := sum([]int{1, 2, 3, 4, 5}, c)
        if err != nil {
            errC <- err // 오류가 발생하면 errC 채널에 오류 전송
            return
        }
    }()

    select { // 여러 채널 중에서 준비된 채널을 선택하여 데이터를 수신함
        case result := <-c:
            expected := 15
            if result != expected {
                t.Errorf("Expected %v, but got %v", expected, result)
            }
        case err := <-errC: // errC 채널에서 오류 수신
            t.Errorf("Unexpected error: %v", err) // 오류가 발생하면 테스트 실패
    }
}

// 테이블 주도 테스트(Table-Driven Tests) 방식으로 테스트 코드 작성
func TestRunConcurrency(t *testing.T) {
    tests := []struct {
        input    []int // 입력값
        expected int // 기대값
        hasError bool
    }{
        {
            []int{7, 2, 8, 4, 0}
            , 21
            , false
        },
        {
            []int{}
            , 0
            , true
        },
    }

    for _, test := range tests {
        result, err := RunConcurrency(test.input)
        if err != nil && !test.hasError { // 오류 케이스 검증이 아닌 경우
            t.Errorf("Unexpected error: %v", err)
        }
        if err == nil && test.hasError { // 오류 케이스 검증
            t.Errorf("Expected error, but got nil")
        }
        if result != test.expected { // 결과 기대값이 다른 경우 검증
            t.Errorf("Expected %v, but got %v", test.expected, result)
        }
    }
}