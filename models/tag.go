package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Tag des
type Tag struct {
	Model
	Name       string `gorm:"type:varchar(64);"json:"name"`
	CreatedBy  int    `gorm:"_"created_by"`
	ModifiedBy int    `gorm:"_"json:"modified_by"`
	State      int    `gorm:"_"json:"state"`
}

// GetTags des
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// GetOneTags des
func GetOneTags(maps interface{}) (tag Tag) {
	db.Where(maps).First(&tag)
	return
}

// GetTagTotal des
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// ExistTagByName des
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// ExistTagByID des
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// EditTag des
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

// AddTag des
func AddTag(name string, state int, createdBy int) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

// DelTag
func DelTagByID(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

// BeforeCreate des
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate des
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
