package config

import (
	"fmt"
	"go-gin/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDB() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			Config.DB.User,
			Config.DB.Password,
			Config.DB.Host,
			Config.DB.Port,
			Config.DB.Database),
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = db
}

func migrateDomains() {
	_ = DB.AutoMigrate(
		&entity.User{},
	)
}
