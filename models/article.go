package models

import (
	"time"

	"gorm.io/gorm"
)

// Article 文章结构体
type Article struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"type:varchar(200);not null" json:"title"`
	Content   string         `gorm:"type:text" json:"content"`
	UserID    uint           `gorm:"not null" json:"user_id"`       // 外键
	User      User           `gorm:"foreignKey:UserID" json:"user"` // 关联用户
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除
}
