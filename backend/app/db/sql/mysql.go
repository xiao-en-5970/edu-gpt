package sql

import (
	"fmt"
	"time"

	"github.com/xiao-en-5970/Goodminton/backend/app/global"
	"github.com/xiao-en-5970/Goodminton/backend/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() {
	// MySQL连接配置
	// dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.Cfg.MySQL.Credentials.Username,
		global.Cfg.MySQL.Credentials.Password,
		global.Cfg.MySQL.Host,
		global.Cfg.MySQL.Port,
		global.Cfg.MySQL.Db,
		)
	
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 可以在这里添加GORM配置
		// 例如: Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		global.Logger.Fatalf("Failed to connect to MySQL: %v", err)
		panic(fmt.Errorf("Failed to connect to MySQL: %v", err))
	}
	
	
	// 获取底层SQL DB连接池
	sqlDB, err := db.DB()
	if err != nil {
		global.Logger.Fatalf("Failed to get DB instance: %v", err)
		panic(fmt.Errorf("Failed to connect to MySQL: %v", err))
	}
	
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(global.Cfg.MySQL.MaxIdleConns)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(global.Cfg.MySQL.SetMaxOpenConns)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(global.Cfg.MySQL.ConnMaxLifetime*time.Hour) // 连接最大存活时间
	global.Db = db
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("自动迁移失败: " + err.Error())
	}
	global.Logger.Infoln("MySQL connected successfully")
	
}