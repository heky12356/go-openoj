package models

import (
	"go-openoj/service/internal/utils"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = utils.Getsql()
}
