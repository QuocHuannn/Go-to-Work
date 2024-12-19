package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid; type:varchar(255); not null; index:idx_uuid; unique;"`
	Username string    `json:"username;"`
	IsActive bool      `gorm:"column:is_active; type:boolean; not null; default:true;"`
	Roles    []Role    `gorm:"many2many:go_user_roles;"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
