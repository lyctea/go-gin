package config

import (
	"go-gin/models"
	"log"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate() error {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Article{},
	)

	if err != nil {
		log.Printf("数据库迁移失败: %v", err)
		return err
	}

	log.Println("数据库表结构迁移成功！")
	return nil
}
