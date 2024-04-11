
---
# Application basicgo
---

- 고 놀이터
- 우선 회원계
- 클린 아키텍쳐 설계


---

# 패키지 구조

---

```
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
```


---

# Go Command

---

### build, run

Go 소스 코드 컴파일 실행
```
go build
go build -o `생성 할 실행파일 이름` `컴파일 할 소스 코드 디렉토리`
go build -o basicgo ./cmd/basicgo
```

생성 된 바이너리 파일 실행
```
./basicgo
```

### test

현재 디렉토리와 모든 하위 디렉토리에서 `_test.go` 로 끝나는 파일을 찾아
`Test`로 시작하는 모든 함수를 찾아 테스트 실행 (해당 함수들은 `*testing.T`를 매개변수로 사용)
```
go test ./...
go test example.com/m/service
```

로그와 함께
```
go test -v example.com/m/service
```

커버리지와 함께
```
go test ./... -cover
```

커버리지와 함께 (html 출력)
```
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```


### 기타

버전 확인
```
go version
// go version go1.22.1 darwin/amd64
```


모듈 초기화
- 현재 디렉토리를 `example.com/m` 과 같은 이름으로 모듈 초기화
`go.mod` 파일을 생성하여 모듈의 이름과 의존성을 정의
```
go mod init example.com/m
```


`go.mod` 파일과 `go.sum` 파일을 업데이트
`go.mod` 파일에 명시된 모든 패키지가 최신 상태가 되도록 하고, 더 이상 사용되지 않는 의존성을 제거하며 모듈에 필요한 모든 의존성을 `go.sum` 파일에 추가합니다.
```
go mod tidy
```


`go get` 명령어를 사용하여 패키지 및 의존성을 추가하며 `$GOPATH/src/<import-path>` 경로에 저장됩니다.
- `$GOROOT`: Go SDK 설치 디렉토리
    - ```ex) /usr/local/opt/go/libexec```
- `$GOPATH`: Go 프로젝트 import 위치, 하위에 src, bin, pkg 디렉토리 존재
    - ```ex) /Users/mac/go```
```
echo $GOPATH
```

`go get` 명령어 옵션
> `go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]`
> -d : 설치는 하지 않고 소스 파일만 다운로드합니다.
> -u : 패키지 및 해당 종속성을 업데이트합니다.
> -t : 패키지에 대한 테스트를 빌드하는 데 필요한 패키지도 다운로드합니다.
> -v : 진행 및 디버그 출력
- $GOPATH/bin 디렉토리에 설치됨


### 추가 의존성


#### gotests
- 테스트 스캐폴딩 자동화
```
go get -u github.com/cweill/gotests/...

gotests -all -w ./cmd/basicgo/main.go
// -all: 모든 함수와 메서드에 대한 테스트 생성
// -w: 생성 된 테스트 코드를 파일로 생성 (해당 옵션이 없으면 stdout 으로 출력함)
또는
gotests -all -w ./*/*
```


#### gRPC
- gRPC 를 사용한 엔드포인트 제공
```
go get -u google.golang.org/grpc

go get -u google.golang.org/protobuf/proto
// go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
// brew install protobuf
// libprotoc 26.1


go get google.golang.org/protobuf/cmd/protoc-gen-go
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

// pb.go 파일 생성
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/user.proto
```

protoc는 Protocol Buffers의 컴파일러로 .proto 파일을 다양한 언어의 소스 코드로 변환하는 역할을 합니다.
이렇게 생성된 소스 코드는 gRPC 클라이언트와 서버를 구현하는 데 사용됩니다.


protoc-gen-go: protoc의 플러그인으로, .proto 파일을 Go 언어의 소스 코드로 변환하는 역할을 합니다. 이 플러그인을 사용하면 .proto 파일에 정의된 서비스와 메시지 타입에 대한 Go 인터페이스와 구조체를 생성할 수 있습니다.


`.proto` 파일에 서비스와 메시지 타입을 정의
- `./proto/user.proto` 참고


`protoc`와 `protoc-gen-go`를 사용하여 이를 Go 코드로 변환
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user.proto
```


grpcurl 설치하여 gRPC 서비스 호출 테스트
```
brew install grpcurl

grpcurl -d '{"id": 1}' -plaintext localhost:50051 user.UserService/GetUser
```