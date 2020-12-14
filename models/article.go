package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (a *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (a *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetArticle(id int) (a Article) {
	db.First(&a, id)
	return
}

func GetArticles(offset, limit int, maps map[string]interface{}) (list []Article) {
	db.Preload("Tag").Where(maps).Offset(offset).Limit(limit).Find(&list)
	return
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func EditArticle(id int, data map[string]interface{}) bool {
	db.Model(&Article{}).Where("id = ? ", id).Updates(data)
	return true
}

func DelArticle(id int) bool {
	db.Where("id = ? ", id).Delete(&Article{})
	return true
}

func ExistsArticleByID(id int) bool {
	var a Article
	db.Where("id = ?", id).First(&a)
	if a.ID > 0 {
		return true
	}
	return false
}

func GetArticlesTotal(maps map[string]interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}
