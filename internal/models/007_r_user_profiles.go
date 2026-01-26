package models

import (
	"time"
)

type UserProfile struct {
	UserProfileNo       uint      `gorm:"primaryKey;column:user_profile_no;type:int(11);autoIncrement" json:"user_profile_no"`
	UserProfileType     string    `gorm:"type:enum('workshop','insurer');not null" json:"user_profile_type"`
	UserProfileName     string    `gorm:"type:varchar(255);not null" json:"user_profile_name"`
	UserProfileCityNo   uint      `gorm:"type:int(11);not null;" json:"user_profile_city_no"`
	UserProfileCityName string    `gorm:"type:varchar(255);not null;" json:"user_profile_city_name"`
	UserProfileAddress  string    `gorm:"type:varchar(255);not null;" json:"user_profile_address"`
	UserProfileEmail    string    `gorm:"type:varchar(255);" json:"user_profile_email"`
	UserProfilePhone    string    `gorm:"type:varchar(30);" json:"user_profile_phone"`
	IsLocked            bool      `gorm:"column:is_locked;type:tinyint(1);not null;default:0" json:"is_locked"`
	CreatedAt           time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	CreatedBy           *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	UpdatedAt           time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy      *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	CreatedByUser      *User  `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User  `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
	City               *City  `gorm:"foreignKey:UserProfileCityNo;references:CityNo" json:"city,omitempty"`
	Users              []User `gorm:"foreignKey:UserProfileNo;references:UserProfileNo;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"users,omitempty"`
}

func (UserProfile) TableName() string {
	return "r_user_profiles"
}
