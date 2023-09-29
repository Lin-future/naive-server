package dao

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitSQLite() (err error) {
	DB, err = gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
