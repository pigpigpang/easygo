package internal

import (
	"easygo/internal/entity"
	"time"

	"easygo/internal/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitGorm 初始化数据库引擎
func InitGorm() (*gorm.DB, error) {
	c := config.C.GORM
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	if c.EnableAutoMigrate {
		err = AutoMigrate(db)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

// NewDB 创建 DB 实例
func NewDB() (*gorm.DB, error) {
	c := config.C.DB
	cGorm := config.C.GORM

	db, err := gorm.Open(mysql.Open(c.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if cGorm.Debug {
		db = db.Debug()
	}

	sql, _ := db.DB()
	err = sql.Ping()
	if err != nil {
		return nil, err
	}

	sql.SetMaxIdleConns(cGorm.MaxIdleConns)
	sql.SetMaxOpenConns(cGorm.MaxOpenConns)
	sql.SetConnMaxLifetime(time.Duration(cGorm.MaxLifetime) * time.Second)

	return db, nil
}

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}
