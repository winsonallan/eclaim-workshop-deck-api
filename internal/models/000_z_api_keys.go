package models

import "time"

type APIKey struct {
	ApiKeyNo    uint       `gorm:"primaryKey;column:api_key_no;type:int(11);autoIncrement" json:"api_key_no"`
	Key         string     `gorm:"column:api_key;type:varchar(64);uniqueIndex;not null" json:"api_key"`
	Name        string     `gorm:"column:key_name;type:varchar(100);not null" json:"key_name"`
	Description string     `gorm:"column:description;type:text" json:"description"`
	IsActive    bool       `gorm:"column:is_active;type:tinyint(1);not null;default:1" json:"is_active"`
	CreatedAt   time.Time  `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	ExpiresAt   *time.Time `gorm:"column:expires_at;type:datetime" json:"expires_at"`
}

func (APIKey) TableName() string {
	return "z_api_keys"
}
