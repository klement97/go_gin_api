package model

import (
	"gorm.io/gorm"
	"italgold/database"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text;" json:"content"`
	UserID  uint
}

func (e *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&e).Error
	if err != nil {
		return &Entry{}, err
	}
	return e, nil
}
