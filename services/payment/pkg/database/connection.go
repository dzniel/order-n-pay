package database

import (
	"fmt"
	"log"

	"github.com/dzniel/order-n-pay/services/payment/pkg/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

type ConnectionInterface interface {
	DBConnect(config helper.Config) (db *gorm.DB, err error)
}

func New() ConnectionInterface {
	return &Connection{
		db: &gorm.DB{},
	}
}

func (conn *Connection) DBConnect(config helper.Config) (db *gorm.DB, err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		config.User, config.Password, config.DBHost, config.DBPort, config.DBName,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection failed:", err)
		return db, fmt.Errorf("database connection failed: " + err.Error())
	}

	log.Println("database connected")
	return db, nil
}
