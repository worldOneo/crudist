package models

import "time"

// GormBaseModel for basic data fields
type GormBaseModel struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GormUser for demonstration
type GormUser struct {
	GormBaseModel
	Username  string     `json:"username,omitempty" gorm:"size:100"`
	Password  string     `json:"password,omitempty" gorm:"size:128"`
	GormPosts []GormPost `json:"gorm_posts"`
}

// GormPost dummy moddel to demonstrate one to many
// in gorm
type GormPost struct {
	GormBaseModel
	GormUserID uint64   `json:"gorm_user_id,omitempty"`
	GormUser   GormUser `json:"gorm_user,omitempty"`
	Title      string   `json:"title,omitempty" gorm:"size:200"`
	Content    string   `json:"content,omitempty" gorm:"size:10000"`
}
