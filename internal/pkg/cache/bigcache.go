package cache

import (
	"github.com/allegro/bigcache"
	"time"
)

var (
	UploadCache *LocalCache
	QueryCache  *LocalCache
)

func init() {
	UploadCache = NewLocalCache()
	QueryCache = NewLocalCache()
}

//修改原来的设置不走定时清理，自动清理会增加一个go协程，定时判断所有的key,因为其时间的判断是基于写入的时间，并且写入的时候会不断往前，所以，这里没必要过多走定时更新
//bigcache在每次set的时候会自动判断最旧的一个是否过期
//当前设置的每个分片最大(256M/64),每个分片初始化大小为 600*512/64*256
func newBigCache() (*bigcache.BigCache, error) {
	cache, err := bigcache.NewBigCache(bigcache.Config{
		CleanWindow:        time.Second * 600,       // 检查Key有效性的时间间隔
		Shards:             64,                      // 最大64分片
		LifeWindow:         time.Second * 3600 * 24, // 60分钟超时
		MaxEntriesInWindow: 100 * 512,               //
		MaxEntrySize:       512,
		HardMaxCacheSize:   256,  // 总共多少内存，这里是256M
		Verbose:            true, // 打印日志
	})

	if err != nil {
		return nil, err
	}

	return cache, nil
}

type LocalCache struct {
	cacheItem *bigcache.BigCache
}

func NewLocalCache() *LocalCache {
	var cache LocalCache
	cache.cacheItem, _ = newBigCache()
	return &cache
}

func (c *LocalCache) Set(key string) {
	if c.cacheItem != nil {
		c.cacheItem.Set(key, []byte("1"))
	}
}

func (c *LocalCache) Del(key string) {
	if c.cacheItem != nil {
		c.cacheItem.Delete(key)
	}
}

func (c *LocalCache) IsExist(key string) (flag bool) {
	if c.cacheItem != nil {
		_, err := c.cacheItem.Get(key)
		if err == nil {
			flag = true
		}
	}
	return
}
