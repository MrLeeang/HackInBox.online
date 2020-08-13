package models

import "time"

// 用户
type UserModel struct {
	Id          int
	Uuid        string
	Name        string
	DisplayName string
	Email       string
	Password    string
	TeamUuid    string
	CreatedAt   time.Time `xorm:"created"`
}
