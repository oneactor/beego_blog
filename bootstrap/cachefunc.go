package bootstrap

import (
	"time"
)

var (
	blogPrefix = "blog_"
)

func (bt *Bootstrap) BlogCachePut(key string, val interface{}) error {
	return Cache.Put(blogPrefix+key, val, blogCacheExpire*time.Second)
}

func (bt *Bootstrap) BlogCacheGet(key string) interface{} {
	return Cache.Get(blogPrefix + key)
}

func (bt *Bootstrap) BlogCacheDelete(key string) error {
	return Cache.Delete(blogPrefix + key)
}

func (bt *Bootstrap) BlogCacheIsExist(key string) bool {
	return Cache.IsExist(blogPrefix + key)
}
