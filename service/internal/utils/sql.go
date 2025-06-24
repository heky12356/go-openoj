package utils

import (
	"fmt"

	"go-openoj/service/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetConfig().Database.User,
		config.GetConfig().Database.Password,
		config.GetConfig().Database.Host,
		config.GetConfig().Database.Name,
	)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Getsql() *gorm.DB {
	return db
}
