package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type MyFormatter struct{}

// 设置日志格式
func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	// 日志格式
	msg := fmt.Sprintf("%s [%s] %s:%v	%s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Caller.File, entry.Caller.Line, entry.Message)
	return []byte(msg), nil
}

// Logger 定义日志
func Logger() *logrus.Logger {
	// now := time.Now()
	// logFilePath := ""
	// if dir, err := os.Getwd(); err == nil {
	// 	logFilePath = dir + "/logs/"
	// }
	// if err := os.MkdirAll(logFilePath, 0777); err != nil {
	// 	fmt.Println(err.Error())
	// }
	// logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	// fileName := path.Join(logFilePath, logFileName)
	// if _, err := os.Stat(fileName); err != nil {
	// 	if _, err := os.Create(fileName); err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }
	//写入文件
	// src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }

	//实例化
	logger := logrus.New()

	//设置输出
	// logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 日志输出
	logger.SetOutput(os.Stdout)

	// 显示行号
	logger.SetReportCaller(true)

	// 设置日志格式
	logger.SetFormatter(new(MyFormatter))

	return logger
}

// UtilsLogger 打印日志调用
var UtilsLogger = Logger()
