package models

import "time"

type Client struct {
	ClientNo            uint      `gorm:"type:int(11);primaryKey;not null;autoincrement" json:"client_no"`
	ClientName          string    `gorm:"type:varchar(100);not null" json:"client_name"`
	ClientEmail         string    `gorm:"type:varchar(255);null" json:"client_email"`
	ClientPhone         string    `gorm:"type:varchar(30);not null" json:"client_phone"`
	Address             string    `gorm:"type:varchar(255);not null" json:"address"`
	CityNo              uint      `gorm:"type:int(11);not null" json:"city_no"`
	CityType            string    `gorm:"type:enum('KAB','KOTA');not null" json:"city_type"`
	CityName            string    `gorm:"type:varchar(255);not null" json:"city_name"`
	VehicleBrandName    string    `gorm:"type:varchar(255);not null" json:"vehicle_brand_name"`
	VehicleSeriesName   string    `gorm:"type:varchar(255);not null" json:"vehicle_series_name"`
	VehicleChassisNo    string    `gorm:"type:varchar(50);not null" json:"vehicle_chassis_no"`
	VehicleLicensePlate string    `gorm:"type:varchar(15);not null" json:"vehicle_license_plate"`
	VehiclePrice        uint      `gorm:"type:bigint(20);not null" json:"vehicle_price"`
	IsLocked            bool      `gorm:"column:is_locked;type:tinyint(1);not null;default:0" json:"is_locked"`
	CreatedBy           *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	CreatedAt           time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	LastModifiedBy      *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`
	UpdatedAt           time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`

	City               *City `gorm:"foreignKey:CityNo;references:CityNo" json:"city,omitempty"`
	CreatedByUser      *User `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (Client) TableName() string {
	return "r_clients"
}
