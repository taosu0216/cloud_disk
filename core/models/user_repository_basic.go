package models

import "time"

type UserRepository struct {
	Id                 int
	ParentId           int64
	UserIdentity       string
	Identity           string
	RepositoryIdentity string
	Ext                string
	Name               string
	CreatedAt          time.Time `xorm:"created_at"`
	UpdatedAt          time.Time `xorm:"updated_at"`
	DeleteAt           time.Time `xorm:"deleted_at"`
}

func (table *UserRepository) TableName() string {
	return "user_repository"
}
