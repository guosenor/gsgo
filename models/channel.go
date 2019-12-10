package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Channel struct {
	Model
	Name     string `gorm:"type:varchar(45);" json:"name"`
	CreateBy int    `gorm:"type:int" json:"createById"`
}

func GetChannels() (channels []Channel) {
	db.Find(&channels)
	return
}

// BeforeCreate des
func (channel *Channel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate des
func (channel *Channel) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
