package db

import (
	"HackInBox.online/models"
	"HackInBox.online/utils"
	"fmt"
)

// QueryUsers 查询所有用户
func QueryUsers() []map[string]interface{} {
	//results, err := Engine.QueryString("select * from user")
	userModel := new(models.UserModel)
	//results, err := Engine.Get(userModel)
	//results, err := Engine.Asc("id").Find(userModel)
	results, err := Engine.Asc("Id").Get(userModel)
	fmt.Println(results)
	fmt.Println(userModel)

	if err != nil {
		utils.Utilslogger.Error(err)
		return []map[string]interface{}{}
	}

	//return results
	return []map[string]interface{}{}
}
