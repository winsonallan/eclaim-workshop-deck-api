package auth

import (
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey;column:user_no;type:int;size:11" json:"user_no"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	Name      string         `gorm:"column:user_name" json:"user_name"`
	IsLocked bool `gorm:"not null;default:0" json:"is_locked"`
	CreatedAt time.Time      `gorm:"column:created_date" json:"created_date"`
	UpdatedAt time.Time      `gorm:"column:last_modified_date" json:"last_modified_date"`
}

func (User) TableName() string {
	return "s_users"
}