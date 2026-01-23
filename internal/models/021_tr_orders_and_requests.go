package models

import "time"

type OrderAndRequest struct {
	OrderRequestNo  uint      `gorm:"primaryKey;type:int(11);autoIncrement;not null" json:"order_request_no"`
	RepairHistoryNo uint      `gorm:"not null;type:int(11)" json:"repair_history_no"`
	SparePartStatus string    `gorm:"type:enum('pending_response','confirmed','shipping','received');not null" json:"spare_part_status"`
	NeededQty       uint      `gorm:"type:int(11);not null" json:"needed_qty"`
	Description     string    `gorm:"not null" json:"description"`
	CreatedAt       time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	CreatedBy       *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	UpdatedAt       time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy  *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	RepairHistory      *RepairHistory `gorm:"foreignKey:RepairHistoryNo;references:RepairHistoryNo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedByUser      *User          `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User          `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (OrderAndRequest) TableName() string {
	return "tr_order_and_requests"
}
