package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

)

type ConnectParams struct {
  Host     string
  User     string
  Password string
  Name     string
  Port     string
}

var DB *gorm.DB

func Connect(params ConnectParams) {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Saigon",
		params.Host,
		params.User,
		params.Password,
		params.Name,
		params.Port,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 10, // Slow SQL threshold
			LogLevel:                  logger.Info,           // Log level
			IgnoreRecordNotFoundError: false,                 // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                  // Disable color
		},
	)

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("database err: ", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("database err: ", err)
	}

	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func Disconnect() {
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("database err: ", err)
	}

	sqlDB.Close()
}
