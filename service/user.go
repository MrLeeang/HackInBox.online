package service

import (
	"HackInBox.online/db"
	"HackInBox.online/models"
)

func GetUsers() []models.UserModel {
	return db.QueryUsers()
}

func GetUserByUuid(uuid string) models.UserModel {
	return db.GetUserByUuid(uuid)
}

func GetUserByTeamUuid(teamUuid string) []models.UserModel {
	return db.DetailUsers("team_uuid", teamUuid)
}

func GetUserByEmail(email string) []models.UserModel {
	return db.DetailUsers("email", email)
}
