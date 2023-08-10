package datalayer

import (
	_ "Kiwi/interfaces"

	"gorm.io/gorm"
)

type kiwiDatalayer struct {
	db *gorm.DB
}
