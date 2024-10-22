package db

import (
	"fmt"
	"time"

	"github.com/AnnonaOrg/osenv"
	"github.com/umfaka/umfaka_core/internal/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// MysqlInit 数据库初始化
func openMysq() (*gorm.DB, error) {
	user := osenv.GetServerDbUsername() //  viper.GetString("db.username")
	pass := osenv.GetServerDbPassword() // viper.GetString("db.password")
	host := osenv.GetServerDbHost()     // viper.GetString("db.host")
	port := osenv.GetServerDbPort()     // viper.GetString("db.port")
	dbname := osenv.GetServerDbName()   // viper.GetString("db.name")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   osenv.GetServerDbTablePrefix(),
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, fmt.Errorf("[db] Database connection failed:(%s) %v", dsn, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("[db] mysql get DB,err: %v", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(1)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(50)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("[db] mysql connDB err: %v", err)
	}
	log.Debug("[db] mysql connDB success")
	return db, nil
}
