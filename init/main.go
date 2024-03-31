package main

import ( // Go 는 사용하지 않는 import 를 허용하지 않음
	"database/sql"
	"log"
	_"github.com/go-sql-driver/mysql" // MySQL driver, 직접 사용되지 않으므로 _ 사용하여 import
	"example.com/m/network"
    "example.com/m/repository"
    "example.com/m/service"
)

func main() {
    // 데이터베이스 연결 설정
    db, err := sql.Open("mysql", "root:dev00@tcp(localhost:3306)/test")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
	
	// UserRepoository, UserService 구현체 생성
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	
	// Server 구현체 생성
	server := network.NewServer(userService, productService)

	// Server 시작
	network.StartServer(server)
}