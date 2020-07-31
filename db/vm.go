package db

import "HackInBox.online/utils"

// QueryVMS 查询所有虚拟机
func QueryVMS() []map[string]interface{} {
	sql2_1 := "select * from vm_ins"
	results, err := Engine.QueryString(sql2_1)
	if err != nil {
		utils.Utilslogger.Error(err)
		results := []map[string]interface{}{}
	}

	return results
}
