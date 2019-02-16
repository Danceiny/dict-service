package service

import (
	"fmt"
	"github.com/Danceiny/go.fastjson"
	"time"
)

import (
	. "github.com/Danceiny/dict-service/persistence"
	. "github.com/Danceiny/dict-service/persistence/entity"
)

var (
	BaseCacheServiceImplCpt *BaseCacheServiceImpl
	ENTITY_CACHE_EXPIRATION time.Duration
	EMPTY_ENTITY_JSON       string
	EMPTY_ENTITY_JSON_BYTES []byte
)

func init() {
	BaseCacheServiceImplCpt = &BaseCacheServiceImpl{RedisImplCpt}
	ENTITY_CACHE_EXPIRATION = 30 * 24 * time.Hour
	EMPTY_ENTITY_JSON = "{\"bid\":null}"
	EMPTY_ENTITY_JSON_BYTES = []byte(EMPTY_ENTITY_JSON)
}

type BaseCacheServiceImpl struct {
	redis Redis
}

func (impl *BaseCacheServiceImpl) CacheEntity(t DictTypeEnum, entity EntityIfc, simple bool) {
	var id = entity.GetBid()
	if t.UseHashCache() {
		impl.redis.HSet(impl.GetTableKey(t), id.String(), fastjson.ToJSON(entity))
	} else {
		impl.redis.Set(impl.GetEntityKey(t, id, simple), fastjson.ToJSON(entity), ENTITY_CACHE_EXPIRATION)
	}

}

func (impl *BaseCacheServiceImpl) MultiCacheEntity(t DictTypeEnum, entities []EntityIfc, simple bool) {
	if entities == nil || len(entities) == 0 {
		return
	}
	var m = make(map[string]interface{})
	if t.UseHashCache() {
		for _, entity := range entities {
			m[entity.GetBid().String()] = fastjson.ToJSON(entity)
		}
		impl.redis.HMSet(impl.GetTableKey(t), m)
	} else {
		for _, entity := range entities {
			m[impl.GetEntityKey(t, entity.GetBid(), simple)] = fastjson.ToJSON(entity)
		}
		impl.redis.MSet(m)
	}
}

func (impl *BaseCacheServiceImpl) CacheEmptyEntity(t DictTypeEnum, bid BID) {
	if t.UseHashCache() {
		impl.redis.HSet(impl.GetTableKey(t), bid.String(), EMPTY_ENTITY_JSON_BYTES)
	} else {
		impl.redis.Pipelined(func(pipe Pipeliner) error {
			pipe.Set(impl.GetEntityKey(t, bid, false), EMPTY_ENTITY_JSON_BYTES, ENTITY_CACHE_EXPIRATION)
			pipe.Set(impl.GetEntityKey(t, bid, true), EMPTY_ENTITY_JSON_BYTES, ENTITY_CACHE_EXPIRATION)
			return nil
		})
	}
}
func (impl *BaseCacheServiceImpl) DeleteEntityCache(t DictTypeEnum, bid BID, simple bool) {
	if t.UseHashCache() {
		impl.redis.HDel(impl.GetTableKey(t), bid.String())
	} else {
		impl.redis.Del(impl.GetEntityKey(t, bid, simple))
	}
}
func (impl *BaseCacheServiceImpl) MultiDeleteEntityCache(t DictTypeEnum, bids []BID, simple bool) {
	size := len(bids)
	keys := make([]string, size)
	if t.UseHashCache() {
		for i, bid := range bids {
			keys[i] = bid.String()
		}
		impl.redis.HDel(impl.GetTableKey(t), keys...)
	} else {
		for i, bid := range bids {
			keys[i] = impl.GetEntityKey(t, bid, simple)
		}
		impl.redis.Del(keys...)
	}
}

func (impl *BaseCacheServiceImpl) GetEntityCache(t DictTypeEnum, bid BID, simple bool) EntityIfc {
	var bytes []byte
	if t.UseHashCache() {
		bytes = RedisImplCpt.HGet(impl.GetTableKey(t), fmt.Sprint(bid))
	} else {
		bytes = RedisImplCpt.Get(impl.GetEntityKey(t, bid, simple))
	}
	return ParseEntityFromJSON(t, bytes)
}

func (impl *BaseCacheServiceImpl) GetEntityKey(t DictTypeEnum, bid BID, simple bool) string {
	var simpleStr string
	if simple {
		simpleStr = "SIMPLE"
	} else {
		simpleStr = "FULL"
	}
	return fmt.Sprintf("Dict:EntityIfc:%s:%s:%v", t, simpleStr, bid)
}

func (impl *BaseCacheServiceImpl) GetTableKey(t DictTypeEnum) string {
	return fmt.Sprintf("Dict:EntityHash:%s", t.String())
}
