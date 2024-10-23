package request

// 信息请求
type BlockedUserRequest struct {
	// ID uuid.UUID `json:"id" form:"id" gorm:"-"`

	Status int `json:"status" form:"status" gorm:"-"`

	UserID    int64  `json:"user_id" form:"user_id" gorm:"-"`
	UserName  string `json:"user_name" form:"user_name" gorm:"-"`
	FirstName string `json:"first_name" form:"first_name" gorm:"-"`
	LastName  string `json:"last_name" form:"last_name" gorm:"-"`

	GroupID int64 `json:"group_id" form:"group_id" gorm:"-"`
	// 机器人id
	BotID int64 `json:"bot_id" form:"bot_id" gorm:"-"`

	// 核验请求id
	ById string `json:"by_id" form:"by_id" gorm:"-"`
	// Boss ID
	BossID int64 `json:"-" form:"-" gorm:"-"`

	Page   int    `json:"-" form:"page" gorm:"-"`
	Size   int    `json:"-" form:"size" gorm:"-"`
	Filter string `json:"-" form:"filter" gorm:"-"`
}
