package experiment

import (
	"fmt"
)

// 슬라이스 합 계산
// a: slice, c: 합산 결과를 전송할 채널
func sum(a []int, c chan int) error {
    total := 0
	if len(a) == 0 {
		return fmt.Errorf("sum Error: Slice is Empty.")
	}
    for _, v := range a {
        total += v
    }
    c <- total // 채널 c에 total 값 전송
	return nil
}

func RunConcurrency(a []int) (int, error) {
    /**
	* 'chan' 채널 Type. 고루틴 간 데이터를 안전하게 전송할 수 있는 통신 케머니즘
	*	- 채널은 버퍼(선택적)를 가질 수 있으며, 버퍼가 가득 차거나 비어있을 때 특정 동작을 수행함
	* 	- 버퍼가 없는 채널: "동기 채널". 채널에 데이터를 보내는 고루틴이 데이터를 받는 고루틴이 준비될 때까지 대기하게 만듦
	* 	- 버퍼가 있는 채널: "비동기 채널". 채널에 데이터를 보내는 고루틴이 버퍼가 가득 찰 때까지 대기하지 않고 즉시 반환하게 만듭니다.
	*/
    c := make(chan int) // make 함수로 버퍼가 없는 채널 c 생성
	errC := make(chan error) // 에러를 전달할 채널 생성

	/**
	* 'go' 키워드로 고루틴 생성
	* 	- 고루틴은 경량 스레드로 함수를 비동기적으로 실행하게 함
	* 	- 메인 고루틴(일반적으로 main 함수)과 다른 고루틴은 서로 블록시키지 않고 계속 실행됨
	* 	- 슬라이스 a의 앞부분과 뒷부분을 각각 다른 고루틴에서 처리 (동시에 처리)
	*/
	go func() {
		errC <- sum(a[:len(a)/2], c)
	}()
	go func() {
		errC <- sum(a[len(a)/2:], c)
	}()

	/**
	* '<-' 연산자를 사용해 채널에 데이터를 보내거나 받을 수 있음
	* 	- 채널에 값이 들어올 때까지 대기하며 먼저 받아온 값이 x 에 할당되고 이후 들어온 값이 y 에 할당됨
	*/
	// x, y := <-c, <-c // 채널 c로부터 값을 받아옴
	// fmt.Println(x, y, x+y)

	var results [2]int
	for i := 0; i < 2; i++ {
		select {
			case err := <-errC:
				if err != nil {
					fmt.Println("Error Response:", err)
					return 0, err
				}
			case res := <-c:
				fmt.Println("Response:", res)
				results[i] = res
		}
	}
	return results[0] + results[1], nil
}