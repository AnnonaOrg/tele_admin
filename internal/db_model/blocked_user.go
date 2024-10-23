package db_model

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

// 黑名单
type BlockedUser struct {
	Model
	Status int `json:"status" form:"status" gorm:"column:status;default:0;not null"`

	UserID    int64  `json:"user_id" form:"user_id" gorm:"column:user_id;index;not null"`
	UserName  string `json:"user_name" form:"user_name" gorm:"column:user_name;"`
	FirstName string `json:"first_name" form:"first_name" gorm:"column:first_name;"`
	LastName  string `json:"last_name" form:"last_name" gorm:"column:last_name;"`

	GroupID int64 `json:"group_id" form:"group_id" gorm:"column:group_id;"`
	// 机器人id
	BotID int64 `json:"bot_id" form:"bot_id" gorm:"column:bot_id;"`
}

func (*BlockedUser) TableName() string {
	return "blocked_user"
}
func (t *BlockedUser) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()

	return
}
func (*BlockedUser) DefaultOrder() string {
	return "updated_at DESC" //updated_at "created_at DESC"
}
func NewBlockedUser(
	userID int64, userName, firstName, lastName string,
	groupID int64,
	botID int64,
	bossID int64,
) *BlockedUser {
	item := &BlockedUser{
		UserID:    userID,
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
		GroupID:   groupID,
		BotID:     botID,
	}
	item.BossID = bossID
	return item
}
