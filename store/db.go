package store

import (
	"github.com/ldigit/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func setup() {
	cfg, _ := config.GetGlobalConfig().(*config.Config)
	dbPath := cfg.GetString("database.path", "reminder.db")
	conn, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Print(err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		log.Printf("connect db server failed. error: %s", err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 600)
	db = conn
}

func NewClient() *gorm.DB {
	if db == nil {
		setup()
	}
	sqlDB, err := db.DB()
	if err != nil {
		setup()
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		setup()
	}

	return db
}
