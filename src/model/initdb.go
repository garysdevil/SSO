package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func OpenDb(username, password, addr, name string) *gorm.DB {
	db, err := gorm.Open("mysql", username+":"+password+"@tcp("+addr+")/"+name+"?charset=utf8&parseTime=true&&loc=Local")

	if err != nil {
		log.Error(err, ",Database connection failed. Database name: ", name)
		panic(err)
	}
	db.LogMode(true)
	log.Info("Database:" + name + " connected")

	return db
}

func InitDB() {

	DB = OpenDb(viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.addr"),
		viper.GetString("db.name"))
	//DB.AutoMigrate(&User{},&Role{})
}
