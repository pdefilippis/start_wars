package datastore

import "gorm.io/gorm"

type IDataStore interface {
	Migrate() error
	GetDB() *gorm.DB
}
