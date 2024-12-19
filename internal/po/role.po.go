package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id;type:int;not null;primaryKey;autoIncrement;comment:'primary key is ID'"`
	Rolename string `gorm:"column:role_name;type:varchar(255);not null"`
	RoleNote string `gorm:"column:role_note;type:text"`
}

func (Role) TableName() string {
	return "go_db_role"
}
