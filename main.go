package main

import (
	"naive-server/dao"
	"naive-server/model"
	"naive-server/route"

	"github.com/sirupsen/logrus"
)

func main() {
	err := dao.InitSQLite()
	if err != nil {
		logrus.Error(err)
		return
	}
	// Migrate the schema
	err = dao.DB.AutoMigrate(&model.User{})
	if err != nil {
		logrus.Error(err)
		return
	}

	route.SetupRouter()
}
