package models

import "time"

type WorkOrder struct {
	WorkOrderNo              uint      `gorm:"type:int(11);primaryKey;not null;autoIncrement" json:"work_order_no"`
	OrderNo                  uint      `gorm:"type:int(11);not null" json:"order_no"`
	AdditionalWorkOrderCount uint      `gorm:"type:tinyint(3);default:0;not null" json:"additional_work_order_count"`
	WorkOrderDocumentNumber  string    `gorm:"type:varchar(50);not null" json:"work_order_document_number"`
	WorkOrderUrl             string    `gorm:"type:varchar(255);not null" json:"work_order_url"`
	IsLocked                 bool      `gorm:"type:tinyint(1);default:0;not null" json:"is_locked"`
	CreatedAt                time.Time `gorm:"column:created_date;type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	CreatedBy                *uint     `gorm:"column:created_by;type:int(11);null" json:"created_by"`
	UpdatedAt                time.Time `gorm:"column:last_modified_date;null;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"last_modified_date"`
	LastModifiedBy           *uint     `gorm:"column:last_modified_by;null;type:int(11)" json:"last_modified_by"`

	Order              *Order `gorm:"foreignKey:OrderNo;references:OrderNo" json:"order,omitempty"`
	CreatedByUser      *User  `gorm:"foreignKey:CreatedBy;references:UserNo" json:"created_by_user,omitempty"`
	LastModifiedByUser *User  `gorm:"foreignKey:LastModifiedBy;references:UserNo" json:"last_modified_by_user,omitempty"`
}

func (WorkOrder) TableName() string {
	return "tr_work_orders"
}
