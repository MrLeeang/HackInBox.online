package db

import (
	"HackInBox.online/model"
	"HackInBox.online/utils"
	"fmt"
)

// QueryVMS 查询所有虚拟机
func QueryUsers() []map[string]interface{} {
	//results, err := Engine.QueryString("select * from user")
	user := new(model.User)
	results, err := Engine.Get(user)
	fmt.Println(results)
	if err != nil {
		utils.Utilslogger.Error(err)
		return []map[string]interface{}{}
	}

	//return results
	return []map[string]interface{}{}
}
