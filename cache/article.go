package cache

import "github.com/iiinsomnia/yiigo"

type ArticleCache struct {
	yiigo.Redis
}

func NewArticleCache() *ArticleCache {
	return &ArticleCache{yiigo.Redis{CacheName: "article"}}
}

func (a *ArticleCache) GetArticleDetail(field string, data interface{}) bool {
	return a.Redis.HGet("detail", field, data)
}

func (a *ArticleCache) SetArticleDetail(field string, data interface{}) bool {
	return a.Redis.HSet("detail", field, data)
}

func (a *ArticleCache) DelArticleDetail(field string) bool {
	return a.Redis.HDel("detail", field)
}
