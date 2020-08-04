package db

import (
	"HackInBox.online/models"
	"HackInBox.online/utils"
)

func CreateDb() {
	utils.Utilslogger.Info("create UserModel")
	Engine.Sync2(new(models.UserModel))
	utils.Utilslogger.Info("create TeamModel")
	Engine.Sync2(new(models.TeamModel))
	utils.Utilslogger.Info("create VmModel")
	Engine.Sync2(new(models.VmModel))
}
