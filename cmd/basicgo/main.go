package main

import ( /* Go 는 사용하지 않는 import 를 허용하지 않음 */
	"log"
	"net"

	// "example.com/m/internal/adapter/controller"
	"example.com/m/internal/adapter/repository"
	"example.com/m/internal/application/usecase"
	"example.com/m/pkg/infrastructure/database"
	// "example.com/m/pkg/infrastructure/web"
	"example.com/m/pkg/infrastructure/config"
	"example.com/m/internal/application/usecase/experiment"

	"google.golang.org/grpc"
    pb "example.com/m/proto"
	/**
		Protocol Buffers 파일 컴파일 후 생성된 Go 코드를 참조하는 import 경로
		- pb: "Protocol Buffers" 의 약어
		- 구조화된 데이터를 직렬화하고, 이를 효율적으로 전송하고 저장하기 위한 Google에서 개발한 언어 중립적인 데이터 직렬화 형식
		- '.proto' 확장자를 가진 파일에 정의 후 프로토콜 버퍼 컴파일러를 사용하여 해당 프로토콜 버퍼 정의를 기반으로 언어별 코드 생성
	*/
	"google.golang.org/grpc/reflection"
)

func main() {
	/**
		HTTP 서버
	// TODO 외부 서비스 클라이언트 초기화 - 현재 미사용
	// productClient := client.NewProductClient("http://external-product-service.com")

	// Repoository, Service, UseCase 초기화
	userRepo := repository.NewUserRepository(db)
	// TODO 외부 서비스 초기화 - 현재 미사용
	// productService := service.NewProductService(productClient)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Controller 초기화
	userController := controller.NewUserController(*userUsecase)

	// Server 구조체를 사용하여 서버 설정 및 시작
	server := web.NewServer(userController)

	// 서버 시작
	web.StartServer(server)
	*/

	/**
		gRPC 서버
	*/
	// DB 연결 설정
	dbConfig := database.DBConfig{
		Username: "root",
		Password: "dev00",
		Host:     "localhost",
		Port:     3306,
		DBName:   "test",
	}
	db := database.NewDBConnection(dbConfig)
	defer db.Close()
	// Repository 초기화
	userRepo := repository.NewUserRepository(db)
	// Usecase 초기화
	userUsecase := usecase.NewUserUsecase(userRepo)

	lis, err := net.Listen("tcp", config.ServerPort)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()

    /**
	*	스터디용 임시 코드
	*/
	// 1. 메모리 테스트
	// experiment.RunMemoryTest()
	// 2. 고루틴 테스트
	experiment.RunConcurrencyTest()

    /**
		gRPC 서버에 서비스 등록
		- gRPC에서는 서비스를 정의하고 해당 서비스에 대한 구현체를 제공하는 방식으로 작동
		- 구현체는 서비스 인터페이스를 구현하는 구조체(서비스에서 정의한 모든 메서드 구현)
	*/
	
	pb.RegisterUserServiceServer(grpcServer, userUsecase)

    // Enable reflection
    reflection.Register(grpcServer)

    log.Println("Server is starting on", config.ServerPort)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}