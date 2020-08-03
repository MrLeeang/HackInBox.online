package db

import (
	"fmt"
	"log"

	"HackInBox.online/utils"
	// 数据库驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

const (
	dbType = "mysql"
	dbHost = "192.168.161.21"
	dbPort = "3306"
	dbUser = "root"
	dbPass = "123456"
	dbName = "maxcloud"
)

// CreateEngine 初始化数据库连接池
func CreateEngine() *xorm.Engine {
	var err error
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
	engine, err := xorm.NewEngine(dbType, dbString)
	if err != nil {
		log.Printf("mysql failed, err:%v\n", err)
	}
	// 设置连接池空闲数
	engine.SetMaxIdleConns(20)
	// 最大打开连接数
	engine.SetMaxOpenConns(100)
	engine.SetLogger(utils.Utilslogger)
	engine.ShowSQL(true)
	return engine
}

// Engine 链接
var Engine = CreateEngine()
