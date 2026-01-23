package models

type City struct {
	CityNo     uint   `gorm:"primaryKey;column:city_no;type:int(11);autoIncrement" json:"city_no"`
	ProvinceNo uint   `gorm:"type:int(11);not null" json:"province_no"`
	CityId     string `gorm:"type:varchar(5);not null" json:"city_id"`
	CityType   string `gorm:"type:enum('KAB','KOTA');not null" json:"city_type"`
	CityName   string `gorm:"type:varchar(70);not null" json:"city_name"`
	IsLocked   bool   `gorm:"column:is_locked;type:tinyint(1);not null;default:0" json:"is_locked"`

	Province *Province `gorm:"foreignKey:ProvinceNo;references:ProvinceNo" json:"province,omitempty"`
}

func (City) TableName() string {
	return "r_cities"
}
