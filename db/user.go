package db

import (
	"HackInBox.online/models"
	"HackInBox.online/utils"
	"fmt"
)

// QueryUsers 查询所有用户
func QueryUsers() []*models.UserModel {
	// 定义一个数组存放结构体
	// allUsers := []*models.UserModel{}
	var allUsers []*models.UserModel
	// 查询
	var err = Engine.Find(&allUsers)

	if err != nil {
		utils.Utilslogger.Error(err)
	}

	for _, user := range allUsers {
		fmt.Println(user)
	}

	return allUsers
}
