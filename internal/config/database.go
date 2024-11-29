package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(global_config GlobalConfig) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/finance?charset=utf8mb4&parseTime=True&loc=Local", global_config.MariaDB.User, global_config.MariaDB.Password, global_config.MariaDB.Host, global_config.MariaDB.Port)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connection established")
}
