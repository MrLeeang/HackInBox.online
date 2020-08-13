package db

import (
	"HackInBox.online/models"
	"HackInBox.online/utils"
	"fmt"
)

// QueryUsers 查询所有用户
func QueryUsers() []models.UserModel {
	// 定义一个数组存放结构体
	// allUsers := []*models.UserModel{}
	var allUsers []models.UserModel
	// 查询
	var err = Engine.Find(&allUsers)

	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return allUsers
}

func GetUserByUuid(uuid string) models.UserModel {
	// 定义一个结构体
	var User models.UserModel
	// 查询
	_, err := Engine.Where("uuid = ?", uuid).Get(&User)

	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return User
}

func DetailUsers(key string, value string) []models.UserModel {
	// 定义一个数组存放结构体
	var UserList []models.UserModel
	// 查询
	sqlString := fmt.Sprintf("%s=?", key)
	var err = Engine.Where(sqlString, value).Find(&UserList)

	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return UserList
}
