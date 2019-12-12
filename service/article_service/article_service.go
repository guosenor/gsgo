package article_service

import (
	"encoding/json"
	"gsgo/models"
	"gsgo/pkg/e"
	"gsgo/pkg/redis"
	"strconv"
)

type Article struct {
	ID    int
	Title string
	Body  string
}

func (article *Article) GetCache() *models.Article {
	var cacheArticle *models.Article
	data, _ := redis.Get(e.CACHE_ARTICLE + strconv.Itoa(article.ID))

	if data != nil {
		json.Unmarshal(data, &cacheArticle)
		return cacheArticle
	}
	a := models.GetArticleByID(article.ID)
	redis.Set(e.CACHE_ARTICLE+strconv.Itoa(a.ID), a, 36000)
	return &a
}
