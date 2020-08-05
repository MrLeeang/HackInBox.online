package models

import "time"

// 战队
type TeamModel struct {
	Id               int
	Uuid             string
	Name             string
	DisplayName      string
	LogoPath         string
	OrganizationName string
	Desc             string
	Num              int
	CreatedAt        time.Time `xorm:"created"`
}
