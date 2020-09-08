package model

import (
	// _ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDb(username, password, addr, dbname string) *gorm.DB {

	dsn := username + ":" + password + "@tcp(" + addr + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// db, err := gorm.Open("mysql", username+":"+password+"@tcp("+addr+")/"+dbname+"?charset=utf8&parseTime=true&&loc=Local")  v1

	if err != nil {
		log.Error(err, ",Database connection failed. Database name: ", dbname)
		panic(err)
	}
	// db
	// db.LogMode(true)
	log.Info("Database:" + dbname + " connected")

	return db
}

func InitDB() {

	DB = OpenDb(viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.addr"),
		viper.GetString("db.name"))
}
