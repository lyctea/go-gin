package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Charset  string
}

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	// 数据库配置
	dbConfig := DatabaseConfig{
		Host:     "47.115.221.211",
		Port:     "3306",
		User:     "learn-go-dev",
		Password: "BHG3PeXnZ2cfn8n6", // 请修改为你的MySQL密码
		DBName:   "learn-go-dev",
		Charset:  "utf8mb4",
	}

	// 构建 DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.Charset,
	)

	// 打开数据库连接
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("获取数据库实例失败: %v", err)
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(100)          // 设置数据库的最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可重用的最大时间

	log.Println("数据库连接成功！")
	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
