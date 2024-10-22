package db

import (
	"fmt"

	"github.com/AnnonaOrg/osenv"
	"gorm.io/gorm"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func DBInit() error {
	db, err := InitSelfDB()
	if err != nil {
		return err
	}
	DB = &Database{
		Self: db,
	}

	MdbTableInit()
	return nil
}

func InitSelfDB() (*gorm.DB, error) {
	switch dbtype := osenv.GetServerDbType(); {
	case dbtype == "mysql":
		return openMysq()
	case dbtype == "sqlite":
		return openSQLite()
	case dbtype == "postgres":
		return openPostgreSQL()
	default:
		if len(dbtype) > 0 {
			return nil, fmt.Errorf("配置中数据库类型(%s)不支持", dbtype)
		}
		return nil, fmt.Errorf("配置中数据库类型(%s)未配置", dbtype)
	}
}

func DBClose() error {
	sqlDB, err := DB.Self.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
