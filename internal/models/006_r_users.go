package models

import (
	"time"
)

type User struct {
	UserNo         uint      `gorm:"primaryKey;column:user_no;type:int(11);autoIncrement" json:"user_no"`
	UserProfileNo  *uint     `gorm:"null;column:user_profile_no; type:int(11);" json:"user_profile_no"`
	RoleNo         uint      `gorm:"column:role_no;type:int(11);" json:"role_no"`
	UserName       string    `gorm:"column:user_name;type:varchar(150)" json:"user_name"`
	UserId         string    `gorm:"column:user_id;type:varchar(255)" json:"user_id"`
	Email          string    `gorm:"column:email;type:varchar(100);uniqueIndex;not null" json:"email"`
	Password       string    `gorm:"column:password;type:varchar(255);not null" json:"-"`
	IsLocked       bool      `gorm:"column:is_locked;type:tinyint(1);not null;default:0" json:"is_locked"`
	CreatedBy      *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	CreatedAt      time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedAt      time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	Role               *Role        `gorm:"foreignKey:RoleNo;references:RoleNo" json:"role,omitempty"`
	UserProfile        *UserProfile `gorm:"foreignKey:UserProfileNo;references:UserProfileNo" json:"user_profile,omitempty"`
	CreatedByUser      *User        `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User        `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (User) TableName() string {
	return "r_users"
}
