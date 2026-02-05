package models

import "time"

type Supplier struct {
	SupplierNo      uint      `gorm:"primaryKey;type:int(11);autoIncrement;not null" json:"supplier_no"`
	WorkshopNo      uint      `gorm:"type:int(11);not null" json:"workshop_no"`
	SupplierName    string    `gorm:"type:varchar(255);not null" json:"supplier_name"`
	SupplierAddress string    `gorm:"type:varchar(255);not null" json:"supplier_address"`
	ProvinceNo      uint      `gorm:"type:int(11);not null" json:"province_no"`
	ProvinceName    string    `gorm:"type:varchar(255);not null" json:"province_name"`
	CityNo          uint      `gorm:"type:int(11);not null" json:"city_no"`
	CityType        string    `gorm:"type:enum('KAB','KOTA');not null" json:"city_type"`
	CityName        string    `gorm:"type:varchar(255);not null" json:"city_name"`
	SupplierPhone   string    `gorm:"type:varchar(30);not null" json:"supplier_phone"`
	SupplierEmail   string    `gorm:"type:varchar(255);not null" json:"supplier_email"`
	IsLocked        bool      `gorm:"column:is_locked;type:tinyint(1);not null;default:0" json:"is_locked"`
	CreatedBy       *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	CreatedAt       time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedAt       time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy  *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	Workshop           *UserProfile `gorm:"foreignKey:WorkshopNo;references:UserProfileNo" json:"workshop,omitempty"`
	Province           *Province    `gorm:"foreignKey:ProvinceNo;references:ProvinceNo" json:"province,omitempty"`
	City               *City        `gorm:"foreignKey:CityNo;references:CityNo" json:"city,omitempty"`
	CreatedByUser      *User        `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User        `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (Supplier) TableName() string {
	return "r_suppliers"
}
