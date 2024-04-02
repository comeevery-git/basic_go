package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver, 직접 사용되지 않으므로 _ 사용하여 import
)

// DBConfig 구조체는 데이터베이스 연결 설정을 위한 구성 정보를 저장합니다.
type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

// NewDBConnection 함수는 DBConfig를 받아서 실제 데이터베이스 연결(*sql.DB)을 생성하고 반환합니다.
func NewDBConnection(cfg DBConfig) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// 데이터베이스 연결 상태를 확인하기 위해 Ping을 시도합니다.
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
	return db
}
