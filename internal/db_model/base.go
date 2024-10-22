package db_model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	// // `gorm:"primarykey;autoIncrement" json:"id"`
	// ID        uint64         `gorm:"primarykey;autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	ID uuid.UUID `json:"id" form:"id" gorm:"column:id;type:uuid;primaryKey;"`
	// Boss ID
	BossID int64 `json:"boss_id" form:"boss_id" gorm:"column:boss_id;"`
	// 备注信息
	Note string `json:"note" form:"note" gorm:"column:note;"`
	// 核验请求id
	ByID string `json:"by_id" form:"by_id" gorm:"-"`
}
