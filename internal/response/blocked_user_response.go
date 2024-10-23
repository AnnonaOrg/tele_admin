package response

// 响应信息
type BlockedUserResponse struct {
	// ID uuid.UUID `json:"id" form:"id" gorm:"-"`

	Status int `json:"status" form:"status" gorm:"-"`

	UserID    int64  `json:"user_id" form:"user_id" gorm:"-"`
	UserName  string `json:"user_name" form:"user_name" gorm:"-"`
	FirstName string `json:"first_name" form:"first_name" gorm:"-"`
	LastName  string `json:"last_name" form:"last_name" gorm:"-"`

	GroupID int64 `json:"group_id" form:"group_id" gorm:"-"`
	// // Boss ID
	// BossID int64 `json:"boss_id" form:"boss_id" gorm:"-"`
	// 机器人id
	BotID int64 `json:"bot_id" form:"bot_id" gorm:"-"`
	// 备注信息
	Note string `json:"note" form:"note" gorm:"-"`
}
