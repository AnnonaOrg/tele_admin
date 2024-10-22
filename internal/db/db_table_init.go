package db

import (
	"sync"

	"github.com/umfaka/umfaka_core/internal/db_model"
	"github.com/umfaka/umfaka_core/internal/log"
)

var once sync.Once

// 自动建表
func MdbTableInit() {
	once.Do(func() {
		if err := DB.Self.AutoMigrate(
			&db_model.BlockedUser{},
		); err != nil {
			log.Errorf("[store_db] AutoMigrate ,err: %v", err)
			return
		}

	})
}
