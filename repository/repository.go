package repository

import (
	"md-tnt-mgmt/iface"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) iface.Repository {
	return &repository{
		db: db,
	}
}
