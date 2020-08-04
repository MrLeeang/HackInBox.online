package service

import "HackInBox.online/db"

func GetUsers() []map[string]interface{} {
	return db.QueryUsers()
}
