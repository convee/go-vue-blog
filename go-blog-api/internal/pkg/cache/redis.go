package cache

import (
	"context"
	"github.com/convee/go-blog-api/pkg/encoding"
	"github.com/convee/go-blog-api/pkg/logger"
	"go.uber.org/zap"
	"reflect"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

// redisCache redis cache结构体
type redisCache struct {
	client            *redis.Client
	KeyPrefix         string
	encoding          encoding.Encoding
	DefaultExpireTime time.Duration
}

// NewRedisCache new一个cache, client 参数是可传入的，方便进行单元测试
func NewRedisCache(client *redis.Client, keyPrefix string, encoding encoding.Encoding) Cache {
	return &redisCache{
		client:    client,
		KeyPrefix: keyPrefix,
		encoding:  encoding,
	}
}

func (c *redisCache) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	buf, err := encoding.Marshal(c.encoding, val)
	if err != nil {
		return errors.Wrapf(err, "marshal data err, value is %+v", val)
	}
	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}
	err = c.client.Set(ctx, cacheKey, buf, expiration).Err()
	if err != nil {
		return errors.Wrapf(err, "redis set err: %+v", err)
	}
	return nil
}

func (c *redisCache) SetNX(ctx context.Context, key string, val interface{}, expiration time.Duration) (bool, error) {
	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return false, errors.Wrapf(err, "build cache key err, key is %+v", key)
	}
	b, err := c.client.SetNX(ctx, cacheKey, val, expiration).Result()
	if err != nil {
		return false, errors.Wrapf(err, "redis set err: %+v", err)
	}
	return b, nil
}

func (c *redisCache) Get(ctx context.Context, key string, val interface{}) error {
	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}

	bytes, err := c.client.Get(ctx, cacheKey).Bytes()
	if err != nil {
		if err != redis.Nil {
			return errors.Wrapf(err, "get data error from redis, key is %+v", cacheKey)
		}
	}

	// 防止data为空时，Unmarshal报错
	if string(bytes) == "" {
		return nil
	}
	if string(bytes) == NotFoundPlaceholder {
		return ErrPlaceholder
	}
	err = encoding.Unmarshal(c.encoding, bytes, val)
	if err != nil {
		return errors.Wrapf(err, "unmarshal data error, key=%s, cacheKey=%s type=%v, json is %+v ",
			key, cacheKey, reflect.TypeOf(val), string(bytes))
	}
	return nil
}

func (c *redisCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	// 批量构建cacheKey
	cacheKeys := make([]string, len(keys))
	for index, key := range keys {
		cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
		if err != nil {
			logger.Warn("build_cache_key_err", zap.String("key", key), zap.Error(err))
			continue
		}
		cacheKeys[index] = cacheKey
	}
	err := c.client.Del(ctx, cacheKeys...).Err()
	if err != nil {
		return errors.Wrapf(err, "redis delete error, keys is %+v", keys)
	}
	return nil
}

func (c *redisCache) HSet(ctx context.Context, key string, filed string, val interface{}) error {
	buf, err := encoding.Marshal(c.encoding, val)
	if err != nil {
		return errors.Wrapf(err, "marshal data err, value is %+v", val)
	}

	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}
	err = c.client.HSet(ctx, cacheKey, cacheKey, filed, buf).Err()
	if err != nil {
		return errors.Wrapf(err, "redis hset err: %+v", err)
	}
	return nil
}

func (c *redisCache) HGet(ctx context.Context, key string, field string, val interface{}) error {
	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}

	bytes, err := c.client.HGet(ctx, cacheKey, field).Bytes()
	if err != nil {
		if err != redis.Nil {
			return errors.Wrapf(err, "get data error from redis, key is %+v", cacheKey)
		}
	}

	// 防止data为空时，Unmarshal报错
	if string(bytes) == "" {
		return nil
	}
	if string(bytes) == NotFoundPlaceholder {
		return ErrPlaceholder
	}
	err = encoding.Unmarshal(c.encoding, bytes, val)
	if err != nil {
		return errors.Wrapf(err, "unmarshal data error, key=%s, cacheKey=%s type=%v, json is %+v ",
			key, cacheKey, reflect.TypeOf(val), string(bytes))
	}
	return nil
}

func (c *redisCache) HDel(ctx context.Context, key string, field string) error {

	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key is %+v", key)
	}
	err = c.client.HDel(ctx, cacheKey, field).Err()
	if err != nil {
		return errors.Wrapf(err, "redis delete error, keys is %+v", key)
	}
	return nil
}

func (c *redisCache) SetCacheWithNotFound(ctx context.Context, key string) error {
	return c.client.Set(ctx, key, NotFoundPlaceholder, DefaultNotFoundExpireTime).Err()
}

// BuildCacheKey 构建一个带有前缀的缓存key
func BuildCacheKey(keyPrefix string, key string) (cacheKey string, err error) {
	if key == "" {
		return "", errors.New("[cache] key should not be empty")
	}

	cacheKey = key
	if keyPrefix != "" {
		cacheKey, err = strings.Join([]string{keyPrefix, key}, ":"), nil
	}

	return
}
