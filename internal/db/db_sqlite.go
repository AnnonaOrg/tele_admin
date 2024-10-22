package db

import (
	"fmt"
	"time"

	"github.com/AnnonaOrg/osenv"
	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 数据库初始化
func openSQLite() (*gorm.DB, error) {
	dbFile := "./conf/db.db"
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   osenv.GetServerDbTablePrefix(),
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, fmt.Errorf("[db] Database connection failed:(%s) %v", dbFile, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("[db] sqlite get DB,err: %v", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(1)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(50)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("[db] sqlite connDB err: %v", err)
	}
	log.Debug("[db] sqlite connDB success")
	return db, nil
}
