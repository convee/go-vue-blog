package cache

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type typeInfo struct {
	fetch      func(...uint64) map[uint64]interface{}
	expire     time.Duration
	expireTime time.Time
	updating   int32
}

type NameCache struct {
	mux      sync.RWMutex
	namesMap map[int]map[uint64]interface{}
	typMap   map[int]*typeInfo
}

func NewNameCache() *NameCache {
	return &NameCache{
		namesMap: map[int]map[uint64]interface{}{},
		typMap:   map[int]*typeInfo{},
	}
}

func (c *NameCache) Register(typ int, expire time.Duration, fetch func(...uint64) map[uint64]interface{}) {
	c.namesMap[typ] = map[uint64]interface{}{}
	c.typMap[typ] = &typeInfo{
		fetch:      fetch,
		expire:     expire,
		expireTime: time.Time{},
		updating:   0,
	}
}

func (c *NameCache) GetString(typ int, id uint64) string {
	v, ok := c.Get(typ, id)
	if !ok {
		fmt.Println("NO NAME FOUND for:", typ, id)
		return ""
	}
	return v.(string)
}

func (c *NameCache) Get(typ int, id uint64) (interface{}, bool) {
	c.mux.RLock()
	tt := c.typMap[typ]
	name, ok := c.namesMap[typ][id]
	c.mux.RUnlock()

	if tt.expireTime.Before(time.Now()) && atomic.CompareAndSwapInt32(&tt.updating, 0, 1) {
		go c.refresh(typ, tt)
	}
	if ok {
		return name, true
	}
	resultMap := tt.fetch(id)
	name, ok = resultMap[id]
	if ok {
		c.mux.Lock()
		defer c.mux.Unlock()
		c.namesMap[typ][id] = name
		return name, true
	}
	return nil, false
}

func (c *NameCache) refresh(typ int, tt *typeInfo) {
	defer atomic.CompareAndSwapInt32(&tt.updating, 1, 0)
	resultMap := tt.fetch()
	if len(resultMap) != 0 {
		c.mux.Lock()
		defer c.mux.Unlock()
		c.namesMap[typ] = resultMap
		tt.expireTime = time.Now().Add(tt.expire)
	}
}

//func (c *NameCache) List(typ int, ids []uint64) map[uint64]string {
//	resultMap := map[uint64]string{}
//	for _, id := range ids {
//		name := c.Get(typ, id)
//		resultMap[id] = name
//	}
//	return resultMap
//}
