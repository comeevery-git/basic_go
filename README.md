# basic_go
고 놀이터



# command

## 기타

Go 버전 확인
- go version
```go version go1.22.1 darwin/amd64```


Go 모듈 초기화
현재 디렉토리를 example.com/m 이라는 이름으로 모듈 초기화. `go.mod` 파일을 생성하여 모듈의 이름과 의존성을 정의
- go mod init example.com/m


`go.mod` 파일과 `go.sum` 파일을 업데이트
`go.mod` 파일에 명시된 모든 패키지가 최신 상태가 되도록 하고, 더 이상 사용되지 않는 의존성을 제거하며 모듈에 필요한 모든 의존성을 `go.sum` 파일에 추가합니다.
- go mod tidy


## build, run

Go 소스 코드 컴파일 실행
- go build
- go build -o `생성 할 실행파일 이름` `컴파일 할 소스 코드 디렉토리`
- go build -o basicgo ./init


생성 된 바이너리 파일 실행
- ./basicgo


## test

현재 디렉토리와 모든 하위 디렉토리에서 `_test.go` 로 끝나는 파일을 찾아 `Test`로 시작하는 모든 함수를 찾아 테스트 실행 (해당 함수들은 `*testing.T`를 매개변수로 사용)
- go test ./...
- go test example.com/m/service


로그와 함께
- go test -v example.com/m/service


