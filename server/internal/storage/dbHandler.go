package storage

import "gorm.io/gorm"

type DbHandler struct {
	DB *gorm.DB
}

func NewDBHandler(db *gorm.DB) DbHandler {

	return DbHandler{db}

}
