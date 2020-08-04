package db

import (
	"HackInBox.online/models"
	"HackInBox.online/utils"
	"fmt"
)

// QueryVMS 查询所有虚拟机
func QueryUsers() []map[string]interface{} {
	//results, err := Engine.QueryString("select * from user")
	userModel := new(models.UserModel)
	results, err := Engine.Get(userModel)
	fmt.Println(results)
	if err != nil {
		utils.Utilslogger.Error(err)
		return []map[string]interface{}{}
	}

	//return results
	return []map[string]interface{}{}
}
