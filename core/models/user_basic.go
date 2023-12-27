package models

import "time"

type UserBasic struct {
	Id        uint64
	Identity  string
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
	DeleteAt  time.Time `xorm:"deleted_at"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
