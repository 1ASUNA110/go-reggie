package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// InitDB 初始化数据库连接
func InitDB() (*gorm.DB, error) {

	// 初始化日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// 格式化 MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.mysql.user"),
		viper.GetString("db.mysql.password"),
		viper.GetString("db.mysql.host"),
		viper.GetInt("db.mysql.port"),
		viper.GetString("db.mysql.database"),
	)

	// 使用 GORM 初始化数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用表名复数
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数化
		},

		// 启动日志
		Logger: newLogger,
	})

	// 检查错误
	if err != nil {
		return nil, err
	}

	// 设置数据库连接池
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)                          // 设置连接最大存活时间
	sqlDB.SetMaxOpenConns(viper.GetInt("db.mysql.maxOpenConns")) // 设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(viper.GetInt("db.mysql.maxIdleConns")) // 设置数据库连接池最大空闲连接数

	log.Println("Database connection established")

	return db, nil
}
