package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgresql 驱动
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbDriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	DbPassword := os.Getenv("DB_PASSWORD")

	// 创建数据库连接
	db, err := sql.Open(DbDriver, fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", DbDriver, DbUser, DbPassword, DbHost, DbPort, DbName))

	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open(DbDriver, db)

	if err != nil {
		fmt.Println("Cannot connect to database ", DbDriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", DbDriver)
	}

	DB.AutoMigrate(&User{})

}

// Cleanup, db connect closed after exits
func Cleanup() {
	if DB != nil {
		_ = DB.Close()
	}
}
