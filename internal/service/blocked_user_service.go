package service

import (
	"fmt"

	"github.com/umfaka/umfaka_core/internal/db_data"
	"github.com/umfaka/umfaka_core/internal/request"
	"github.com/umfaka/umfaka_core/internal/response"
)

// func SetBlockedUser(userID int64, botID int64) error {
// 	return db_data.SetBlockedUser(userID, botID)
// }
// func CheckBlockedUser(userID int64, botID int64) bool {
// 	return db_data.CheckBlockedUser(userID, botID)
// }

func CreateBlockedUser(
	userID int64, userName, firstName, lastName string,
	groupID int64,
	botID int64,
	bossID int64,
) (*response.BlockedUserResponse, error) {
	// 如果已经在屏蔽列表，终止后续操作
	if count, _ := GetCountBlockedUserByUserIDAndGroupID(userID, groupID); count > 0 {
		return nil, fmt.Errorf("GetCountBlockedUserByUserIDAndGroupID(%d,%d): %d", userID, groupID, count)
	}
	item, err := db_data.CreateBlockedUser(
		userID, userName, firstName, lastName,
		groupID,
		botID,
		bossID,
	)
	if err != nil {
		return nil, err
	}
	resp := &response.BlockedUserResponse{
		Status: item.Status,

		UserID:    item.UserID,
		UserName:  item.UserName,
		FirstName: item.FirstName,
		LastName:  item.LastName,
		GroupID:   item.GroupID,
		// BossID:    item.BossID,
		BotID: item.BotID,
		Note:  item.Note,
	}
	return resp, nil
}
func CreateBlockedUserEx(req *request.BlockedUserRequest) (*response.BlockedUserResponse, error) {
	return CreateBlockedUser(req.UserID, req.UserName, req.FirstName, req.LastName, req.GroupID, req.BotID, req.BossID)
}
func GetBlockedUserByUsername(username string, groupID int64) (*response.BlockedUserResponse, error) {
	item, err := db_data.GetBlockedUserByUsername(username, groupID)
	if err != nil {
		return nil, err
	}
	resp := &response.BlockedUserResponse{
		Status: item.Status,

		UserID:    item.UserID,
		UserName:  item.UserName,
		FirstName: item.FirstName,
		LastName:  item.LastName,
		GroupID:   item.GroupID,
		// BossID:    item.BossID,
		BotID: item.BotID,
		Note:  item.Note,
	}
	return resp, nil
}

//	func GetCountBlockedUserByUserIDAndBotID(userID, botID int64) (int64, error) {
//		return db_data.GetCountBlockedUserByUserIDAndBotID(userID, botID)
//	}
//
//	func DeleteBlockedUser(userID, botID int64) error {
//		return db_data.DeleteBlockedUser(userID, botID)
//	}
func GetCountBlockedUserByUserIDAndGroupID(userID, groupID int64) (int64, error) {
	return db_data.GetCountBlockedUserByUserIDAndGroupID(userID, groupID)
}
func DeleteBlockedUser(userID, groupID int64) error {
	return db_data.DeleteBlockedUser(userID, groupID)
}
