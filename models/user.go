package models

import "time"

// 用户
type UserModel struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Team_uuid string
	CreatedAt time.Time `xorm:"created"`
}
