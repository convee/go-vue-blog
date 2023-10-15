package cache

import (
	"time"

	"github.com/convee/go-blog-api/pkg/encoding"
	"github.com/convee/go-blog-api/pkg/redis"
)

const (
	RedisKeyCustomer = "customer:"
	RedisTTL         = 600 * time.Second
)

// NewUserCache new一个用户cache
func NewUserCache() Cache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := "ms"
	return NewRedisCache(redis.RedisClient, cachePrefix, jsonEncoding)
}
