package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
        // args []string
	}{
		// TODO: Add test cases.
		{
			name: "TestMain",
		},
		// Add test case Example..
        // {
        //     name: "TestMainWithArgs",
        //     args: []string{"cmd", "arg1", "arg2"},
        // },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
			// TODO: main 실행 후 로그 파일 생성이나 DB 연결이나 데이터 생성 확인 테스트 좋을 듯
		})
	}
}
