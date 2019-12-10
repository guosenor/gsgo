package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	Title     string  `gorm:"type:varchar(45);" json:"title"`
	Body      string  `gorm:"typ:TEXT;" json:"body"`
	ChannelID int     `json:"channelId"`
	Channel   Channel `json:"channel"`
	Tags      []Tag   `gorm:"many2many:article_tags" json:"tags"`
	CreateBy  int     `gorm:"type:int" json:"createById"`
}

// BeforeCreate des
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate des
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}
func GetArticleByID(id int) (article Article) {
	db.First(&article, id)
	db.Model(&article).Related(&article.Channel, "ChannelID")
	return
}

func AddArticle(title string, body string, channelId int, createById int) (article Article) {
	article.ChannelID = channelId
	article.Title = title
	article.Body = body
	article.CreateBy = createById
	fmt.Println(article.ChannelID)
	db.Create(&article)
	db.Model(&article).Related(&article.Channel, "ChannelID")
	return
}
