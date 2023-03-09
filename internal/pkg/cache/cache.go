package cache

import (
	"context"
	"errors"
	"time"
)

var (
	// DefaultExpireTime 默认过期时间
	DefaultExpireTime = time.Hour * 24
	// DefaultNotFoundExpireTime 结果为空时的过期时间 1分钟, 常用于数据为空时的缓存时间(缓存穿透)
	DefaultNotFoundExpireTime = time.Minute
	// NotFoundPlaceholder .
	NotFoundPlaceholder = "*"

	// DefaultClient 生成一个缓存客户端，其中keyPrefix 一般为业务前缀
	DefaultClient Cache

	// ErrPlaceholder .
	ErrPlaceholder = errors.New("cache: placeholder")
	// ErrSetMemoryWithNotFound .
	ErrSetMemoryWithNotFound = errors.New("cache: set memory cache err for not found")
)

// Cache 定义cache驱动接口
type Cache interface {
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	SetNX(ctx context.Context, key string, val interface{}, expiration time.Duration) (bool, error)
	Get(ctx context.Context, key string, val interface{}) error
	Del(ctx context.Context, keys ...string) error
	HSet(ctx context.Context, key string, field string, val interface{}) error
	HGet(ctx context.Context, key string, field string, val interface{}) error
	HDel(ctx context.Context, key string, field string) error
	SetCacheWithNotFound(ctx context.Context, key string) error
}

// Set 数据
func Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	return DefaultClient.Set(ctx, key, val, expiration)
}

// SetNX 数据
func SetNX(ctx context.Context, key string, val interface{}, expiration time.Duration) (bool, error) {
	return DefaultClient.SetNX(ctx, key, val, expiration)
}

// Get 数据
func Get(ctx context.Context, key string, val interface{}) error {
	return DefaultClient.Get(ctx, key, val)
}


// Del 批量删除
func Del(ctx context.Context, keys ...string) error {
	return DefaultClient.Del(ctx, keys...)
}

// HSet 数据
func HSet(ctx context.Context, key string, field string, val interface{}) error {
	return DefaultClient.HSet(ctx, key, field, val)
}

// HGet 数据
func HGet(ctx context.Context, key string, field string, val interface{}) error {
	return DefaultClient.HGet(ctx, key, field, val)
}


// HDel 批量删除
func HDel(ctx context.Context, key string, field string) error {
	return DefaultClient.HDel(ctx, key, field)
}
// SetCacheWithNotFound .
func SetCacheWithNotFound(ctx context.Context, key string) error {
	return DefaultClient.SetCacheWithNotFound(ctx, key)
}
