package article_service

import (
	"encoding/json"
	"gsgo/models"
	"gsgo/pkg/e"
	"gsgo/pkg/redis"
	"strconv"
	"strings"
)

type Article struct {
	ID         int
	Title      string
	Body       string
	CreateByID int
	ChannelID  int
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
func (article *Article) GetArticlesCache() (articles []*models.Article) {
	where := make(map[string]interface{})
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
	}
	if article.ChannelID > 0 {
		keys = append(keys, strconv.Itoa(article.ChannelID))
		where["channel_id"] = article.ChannelID
	}
	if article.CreateByID > 0 {
		keys = append(keys, strconv.Itoa(article.CreateByID))
		where["create_by"] = article.CreateByID
	}
	key := strings.Join(keys, "_")
	data, _ := redis.Get(key)
	if data != nil {
		json.Unmarshal(data, &articles)
		return
	}
	articles = models.GetArticles(where)
	redis.Set(key, articles, 36000)
	return
}
