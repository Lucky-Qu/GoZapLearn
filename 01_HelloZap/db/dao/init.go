package dao

import (
	"GoZapLearn/01_HelloZap/config"
	"GoZapLearn/01_HelloZap/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: config.DSN,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			TablePrefix:   "ZipLearn_",
		},
	})
	if err != nil {
		panic(err)
	}
	if !db.Migrator().HasTable(&model.Person{}) {
		err = db.Migrator().CreateTable(&model.Person{})
		if err != nil {
			panic(err)
		}
	}
}
