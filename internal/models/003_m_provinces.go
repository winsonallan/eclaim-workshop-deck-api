package models

type Province struct {
	ProvinceNo   uint   `gorm:"primaryKey;column:province_no;type:int(11);autoIncrement" json:"province_no"`
	ProvinceId   uint   `gorm:"type:smallint(3);not null;column:province_id" json:"province_id"`
	ProvinceName string `gorm:"type:varchar(50);not null" json:"province_name"`
	IsLocked     bool   `gorm:"column:is_locked;type:tinyint(1);not null;default:0" json:"is_locked"`
}

func (Province) TableName() string {
	return "m_provinces"
}
