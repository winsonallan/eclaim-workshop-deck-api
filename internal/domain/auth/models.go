package auth

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;column:user_no;type:int(11);autoIncrement" json:"user_no"`
	Email     string    `gorm:"column:email;type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"column:password;type:varchar(255);not null" json:"-"`
	Name      string    `gorm:"column:user_name;type:varchar(100)" json:"user_name"`
	IsLocked  bool      `gorm:"column:is_locked;type:tinyint(1);not null;default:0" json:"is_locked"`
	CreatedAt time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedAt time.Time `gorm:"column:last_modified_date;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
}

func (User) TableName() string {
	return "s_users"
}

// API Key Model - ADD THIS
type APIKey struct {
	ID          uint       `gorm:"primaryKey;column:api_key_no;type:int(11);autoIncrement" json:"api_key_no"`
	Key         string     `gorm:"column:api_key;type:varchar(64);uniqueIndex;not null" json:"api_key"`
	Name        string     `gorm:"column:key_name;type:varchar(100);not null" json:"key_name"`
	Description string     `gorm:"column:description;type:text" json:"description"`
	IsActive    bool       `gorm:"column:is_active;type:tinyint(1);not null;default:1" json:"is_active"`
	CreatedAt   time.Time  `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	ExpiresAt   *time.Time `gorm:"column:expires_at;type:datetime" json:"expires_at"`
}

func (APIKey) TableName() string {
	return "s_api_keys"
}