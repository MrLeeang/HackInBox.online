package db

import (
	models "HackInBox.online/models"
	"HackInBox.online/utils"
)

// QueryUsers 查询所有用户
func QueryUsers() []models.UserModel {
	// 定义一个数组存放结构体
	// allUsers := []*models.UserModel{}
	var allUsers []models.UserModel
	// 查询
	var err = Engine.Find(&allUsers)

	if err != nil {
		utils.Utilslogger.Error(err)
	}

	return allUsers
}

func GetUserByUuid(uuid string) models.UserModel {
	// 定义一个结构体
	var User models.UserModel
	// 查询
	_, err := Engine.Where("uuid = ?", uuid).Get(&User)

	if err != nil {
		utils.Utilslogger.Error(err)
	}

	return User
}

//func DetailUsers(uuid string) []models.UserModel {
//	// 定义一个数组存放结构体
//	var User []models.UserModel
//	// 查询
//	var err = Engine.Where("uuid=?", uuid).Find(&User)
//
//	if err != nil {
//		utils.Utilslogger.Error(err)
//	}
//
//	return User
//}
