package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户结构体
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Age       int            `gorm:"not null" json:"age"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除
}
