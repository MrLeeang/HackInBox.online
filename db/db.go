package db

import (
	"HackInBox.online/utils"
	"fmt"
	// 数据库驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

const (
	dbType = "mysql"
	dbHost = "118.31.237.114"
	dbPort = "3307"
	dbUser = "root"
	dbPass = "123456"
	dbName = "HackInBox"
)

// CreateEngine 初始化数据库连接池
func CreateEngine() *xorm.Engine {
	var err error
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
	engine, err := xorm.NewEngine(dbType, dbString)
	if err != nil {
		utils.UtilsLogger.Error(err)
	}
	// 设置连接池空闲数
	engine.SetMaxIdleConns(20)
	// 最大打开连接数
	engine.SetMaxOpenConns(100)
	engine.SetLogger(utils.UtilsLogger)
	return engine
}

// Engine 链接
var Engine = CreateEngine()
