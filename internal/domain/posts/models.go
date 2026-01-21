package posts

import (
	"eclaim-workshop-deck-api/internal/domain/auth"
	"time"
)

type Post struct {
	ID        uint           `gorm:"column:post_no;primaryKey;type:int;size:11" json:"post_no"`
	PostTitle     string         `gorm:"not null" json:"post_title"`
	PostContent   string         `gorm:"type:text" json:"post_content"`
	UserNo    uint           `gorm:"not null;type:int;size:11" json:"user_no"`
	IsLocked bool `gorm:"not null;default:0" json:"is_locked"`
	User      auth.User      `gorm:"foreignKey:UserNo;references:ID; comment:References the primary key of the s_users table" json:"user"`
	CreatedAt time.Time      `gorm:"column:created_date" json:"created_date"`
	UpdatedAt time.Time      `gorm:"column:last_modified_date" json:"last_modified_date"`
}

func (Post) TableName() string {
	return "s_posts"
}