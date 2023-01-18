package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web-study/settings"
)

var db *gorm.DB

func Init(cfg *settings.MySQLConfig) error {
	//init函数是在main函数运行前运行的函数
	//grom 2。0 之后的连接方式
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return nil
	}
	//TODO:将表配置到数据库中去
	if err := db.AutoMigrate(&User{}, &Community{}); err != nil {
		zap.L().Error("AutoMigrate DB failed", zap.Error(err))
		return nil
	}
	a, err := db.DB()
	a.SetMaxOpenConns(cfg.MaxOpenConns)
	a.SetMaxIdleConns(cfg.MaxIdleConns)
	return nil

}

func Close() {
	a, err := db.DB()
	err = a.Close()
	if err != nil {
		zap.L().Error(fmt.Sprintf("close mysql failed,err:%v", err))
	}
}
