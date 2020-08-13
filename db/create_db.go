package db

import (
	"HackInBox.online/models"
	"HackInBox.online/utils"
)

func CreateDb() {
	utils.UtilsLogger.Info("create UserModel")
	Engine.Sync2(new(models.UserModel))
	utils.UtilsLogger.Info("create TeamModel")
	Engine.Sync2(new(models.TeamModel))
	utils.UtilsLogger.Info("create VmModel")
	Engine.Sync2(new(models.VmModel))
}
