package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

type MariaDBConfig struct {
	Address      string
	DatabaseName string
	Username     string
	Password     string
	Port         string
}

func init() {
	config := MariaDBConfig{
		Address:      "db",
		DatabaseName: "mariadb",
		Username:     "mariadb",
		Password:     "mariadb",
		Port:         "3307",
	}
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", config.Username, config.Password, config.Address, config.DatabaseName)

	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
}
