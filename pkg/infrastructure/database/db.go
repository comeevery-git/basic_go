package database

import (
	// "database/sql"
	"fmt"
	"log"

	// _ "github.com/go-sql-driver/mysql" // MySQL driver, 직접 사용되지 않으므로 _ 사용하여 import
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql" // MySQL driver, go-sql-driver/mysql 을 내부적으로 사용함
)

// DBConfig 구조체는 데이터베이스 연결 설정을 위한 구성 정보를 저장합니다.
type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

// NewDBConnection 함수는 DBConfig를 받아서 실제 데이터베이스 연결(*gorm.DB)을 생성하고 반환함
func NewDBConnection(cfg DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// sql.db 연결 상태 확인을 위해 Ping 호출
	// 	- Gorm 의 Open 메서드는 DB 연결 시 내부적으로 Ping 을 호출해서 연결 상태를 확인하기에 해당 부분 불필요
	// if err := db.Ping(); err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }

	log.Println("Database connection established")
	return db
}
