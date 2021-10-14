package models

import "time"

type GormBaseModel struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GormUser struct {
	GormBaseModel
	Username string `json:"username" gorm:"size:100"`
	Password string `json:"password" gorm:"size:128"`
}