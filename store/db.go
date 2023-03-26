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
	if cfg == nil {
		log.Print("GetGlobalConfig failed")
		return
	}
	dsn := cfg.GetString("database.dsn", "reminder.db")
	conn, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		log.Printf("connect db server failed. error: %s", err.Error())
	}
	sqlDB.SetMaxIdleConns(cfg.GetInt("database.max_idle", 10))
	sqlDB.SetMaxOpenConns(cfg.GetInt("database.max_open", 100))
	sqlDB.SetConnMaxLifetime(time.Second * 600)
	db = conn
}

// MigrateTable 初始化数据表
func MigrateTable() {
	db := GetInstance()
	tables := []interface{}{
		Message{},
		Task{},
	}

	for _, t := range tables {
		_ = db.AutoMigrate(&t)
	}
}

// GetInstance 获取DB连接实例
func GetInstance() *gorm.DB {
	if db == nil {
		setup()
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			setup()
		}
		if err = sqlDB.Ping(); err != nil {
			sqlDB.Close()
			setup()
		}
	}

	return db
}
