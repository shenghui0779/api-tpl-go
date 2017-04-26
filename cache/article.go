package cache

import "github.com/iiinsomnia/yiigo"

type ArticleCache struct {
	yiigo.Redis
}

func NewArticleCache() *ArticleCache {
	return &ArticleCache{yiigo.Redis{CacheName: "article"}}
}

func (a *ArticleCache) GetArticle(field string, data interface{}) bool {
	return a.Redis.HGet("detail", field, data)
}

func (a *ArticleCache) SetArticle(field string, data interface{}) bool {
	return a.Redis.HSet("detail", field, data)
}

func (a *ArticleCache) DelArticle(field string) bool {
	return a.Redis.HDel("detail", field)
}
