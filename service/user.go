package service

import (
	"HackInBox.online/db"
	"HackInBox.online/models"
)

func GetUsers() []*models.UserModel {
	return db.QueryUsers()
}
