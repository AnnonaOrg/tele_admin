package db_data

import (
	"fmt"

	"github.com/umfaka/umfaka_core/internal/db"

	"github.com/umfaka/umfaka_core/internal/db_model"
	"github.com/umfaka/umfaka_core/internal/log"
)

func SetBlockedUser(userID int64, botID int64) error {
	key := fmt.Sprintf("Tele_Admin_Black_List_%d", botID)

	return AddToSet(key, userID)
}
func CheckBlockedUser(userID int64, botID int64) bool {
	key := fmt.Sprintf("Tele_Admin_Black_List_%d", botID)

	if isBlocked, err := IsMemberOfSet(key, userID); err != nil {
		log.Errorf("IsMemberOfSet(%s,%d): %v", key, userID, err)
		return false
	} else {
		return isBlocked
	}
}

// 创建信息
func CreateBlockedUser(
	userID int64, userName, firstName, lastName string,
	botID int64,
	bossID int64,
) (*db_model.BlockedUser, error) {
	item := db_model.NewBlockedUser(
		userID, userName, firstName, lastName,
		botID,
		bossID,
	)
	err := db.DB.Self.Create(item).Error
	return item, err
}

func GetBlockedUserByUsername(username string) (*db_model.BlockedUser, error) {
	item := new(db_model.BlockedUser)
	err := db.DB.Self.Model(item).
		Where("user_name = ?", username).
		First(item).
		Error
	return item, err
}

func GetCountBlockedUserByUserIDAndBotID(userID, botID int64) (int64, error) {
	var count int64
	item := new(db_model.BlockedUser)
	err := db.DB.Self.Model(item).
		Where("user_id = ? AND bot_id = ?", userID, botID).
		Count(&count).
		Error
	return count, err
}

// 删除信息记录
func DeleteBlockedUser(userID, botID int64) error {
	result := db.DB.Self.
		Where("bot_id = ? AND user_id = ?", botID, userID).
		Delete(&db_model.BlockedUser{})
	if result.RowsAffected == 0 {
		return fmt.Errorf("Delete(%d,%d): not_found", botID, userID)
	}
	return nil
}
