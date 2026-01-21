package posts

import (
	"eclaim-workshop-deck-api/internal/domain/auth"
	"time"
)

type Post struct {
	PostNo        		uint           `gorm:"column:post_no;primaryKey;type:int(11);autoIncrement" json:"post_no"`
	PostTitle     string         `gorm:"not null;type:varchar(255)" json:"post_title"`
	PostContent   string         `gorm:"type:text" json:"post_content"`
	UserNo    		uint           `gorm:"not null;type:int(11);comment:References the primary key of the s_users table" json:"user_no"`
	IsLocked 			bool 					 `gorm:"not null;default:0" json:"is_locked"`
	User      		auth.User      `gorm:"foreignKey:UserNo;references:UserNo" json:"user"`
	CreatedAt 		time.Time      `gorm:"column:created_date" json:"created_date"`
	UpdatedAt 		time.Time      `gorm:"column:last_modified_date" json:"last_modified_date"`
}

func (Post) TableName() string {
	return "s_posts"
}