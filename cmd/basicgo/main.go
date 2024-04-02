package main

import ( // Go 는 사용하지 않는 import 를 허용하지 않음
	"example.com/m/internal/adapter/controller"
	"example.com/m/internal/adapter/repository"
	"example.com/m/internal/application/usecase"
	"example.com/m/pkg/infrastructure/database"
	"example.com/m/pkg/infrastructure/web"
)

func main() {
    // 데이터베이스 연결 설정
	dbConfig := database.DBConfig{
		Username: "root",
		Password: "dev00",
		Host:     "localhost",
		Port:     3306,
		DBName:   "test",
	}
	db := database.NewDBConnection(dbConfig)
	defer db.Close()

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
	web.StartServer(server)
}