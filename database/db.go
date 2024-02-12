package database

import (
	//"io/fs"
	//"os"
	//"path"
	"x-ui/config"
	"x-ui/database/model"

	"gorm.io/driver/mysql"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func initUser() error {
	err := db.AutoMigrate(&model.V2rayUser{})
	if err != nil {
		return err
	}
	var count int64
	err = db.Model(&model.V2rayUser{}).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		user := &model.V2rayUser{
			Username: "admin",
			Password: "admin",
		}
		return db.Create(user).Error
	}
	return nil
}

func initInbound() error {
	return db.AutoMigrate(&model.V2rayInbound{})
}

func initSetting() error {
	return db.AutoMigrate(&model.V2raySetting{})
}

func InitDB(dbPath string) error {
	/*dir := path.Dir(dbPath)
	err := os.MkdirAll(dir, fs.ModeDir)
	if err != nil {
		return err
	}*/

	var err error
	var gormLogger logger.Interface

	if config.IsDebug() {
		gormLogger = logger.Default
	} else {
		gormLogger = logger.Discard
	}

	c := &gorm.Config{
		Logger: gormLogger,
	}
	db, err = gorm.Open(mysql.Open(config.GetMySqlPath()), c)
	if err != nil {
		return err
	}

	err = initUser()
	if err != nil {
		return err
	}
	err = initInbound()
	if err != nil {
		return err
	}
	err = initSetting()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
