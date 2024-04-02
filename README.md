# Application basicgo
고 놀이터
- 우선 회원계
- 클린 아키텍쳐 설계

# 패키지 구조

.
├── cmd                             # 애플리케이션 진입점. 실행 가능한 바이너리 설정.
│   └── basicgo
│       └── main.go
├── domain                          # 엔터티와 비즈니스 규칙. 애플리케이션의 핵심 로직
│   ├── model                       # [Entity]
│   │   └── model.go                # 비즈니스 엔터티. 데이터와 행위를 캡슐화
│   └── repository
│       └── user.go       # 데이터 저장소 인터페이스
├── external                        # 외부 서비스 및 Core 서버 간 통신 구현
│   ├── client                      # 다른 Core 서버와 통신을 위한 클라이언트 구현체
│   │   └── product.go
│   ├── dto
│   │   └── product.go              # 외부 서비스로부터 받은 데이터의 구조 정의
│   └── service                     # 외부 API 호출 및 통신을 추상화하는 서비스
│       └── product.go
├── go.mod
├── go.sum
├── internal                        # 애플리케이션 내부 로직
│   ├── adapter                     # 인터페이스 어댑터 계층: 데이터 포맷 변환, 외부 통신 중개 등
│   │   ├── controller              # [Adapter] HTTP 요청 처리: 사용자의 요청을 비즈니스 로직(usecase)으로 전달
│   │   │   └── user.go
│   │   ├── presenter               # 비즈니스 로직 결과를 클라이언트에 적합한 형태로 변환
│   │   │   └── user.go
│   │   └── repository              # 데이터 저장소 인터페이스의 구현체
│   │       ├── database.sql        # SQL 쿼리문
│   │       ├── user.go
│   │       └── ...
│   └── application                 # 응용 프로그램 계층. 비즈니스 로직과 유스 케이스 구현
│       └── usecase                 # [UseCase] 핵심 비즈니스 로직을 구현하는 서비스
│           ├── user.go
│           └── ...
├── pkg                             # 재사용 가능한 코드 모듈. 다른 프로젝트에서도 사용할 수 있는 라이브러리 코드
│   ├── infrastructure              # [Frameworks&Drivers] 기술적 세부 사항과 외부 의존성, 메시징 시스템 등
│   │   ├── config
│   │   ├── database                # 데이터베이스 연결 관리
│   │   └── web                     # HTTP 서버 구성 및 요청 라우팅 관리
│   │   │   └── network.go
│   ├── types                       # 애플리케이션 전반에서 사용되는 공통 타입 정의
│   │   └── errors                  # 에러 처리 관련 코드 및 커스텀 에러 타입 정의
│   └── utils                       # 애플리케이션 전반에서 사용되는 공통 기능 정의
│       └── utils.go
└── README.md


# Go Command


## build, run

Go 소스 코드 컴파일 실행
- go build
- go build -o `생성 할 실행파일 이름` `컴파일 할 소스 코드 디렉토리`
- go build -o basicgo ./cmd/basicgo


생성 된 바이너리 파일 실행
- ./basicgo


## test

현재 디렉토리와 모든 하위 디렉토리에서 `_test.go` 로 끝나는 파일을 찾아 `Test`로 시작하는 모든 함수를 찾아 테스트 실행 (해당 함수들은 `*testing.T`를 매개변수로 사용)
- go test ./...
- go test example.com/m/service


로그와 함께
- go test -v example.com/m/service



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


