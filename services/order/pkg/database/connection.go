package database

import (
	"fmt"
	"log"

	"github.com/dzniel/order-n-pay/services/order/pkg/helper"
	"gorm.io/driver/postgres"
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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.Timezone,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection failed:", err)
		return db, fmt.Errorf("database connection failed: " + err.Error())
	}

	log.Println("database connected")
	return db, nil
}
