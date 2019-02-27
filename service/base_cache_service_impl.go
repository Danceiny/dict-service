package service

import (
    "fmt"
    "time"
)

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence"
    . "github.com/Danceiny/dict-service/persistence/entity"
)

var (
    ENTITY_CACHE_EXPIRATION time.Duration
    EMPTY_ENTITY_JSON       string
    EMPTY_ENTITY_JSON_BYTES []byte
)

func init() {
    ENTITY_CACHE_EXPIRATION = 30 * 24 * time.Hour
    EMPTY_ENTITY_JSON = "{\"bid\":null}"
    EMPTY_ENTITY_JSON_BYTES = []byte(EMPTY_ENTITY_JSON)
}

type BaseCacheServiceImpl struct {
    Cache Redis
}

func (impl *BaseCacheServiceImpl) CacheEntity(t DictTypeEnum, entity EntityIfc, simple bool) {
    var id = entity.GetBid()
    if t.UseHashCache() {
        impl.Cache.HSet(impl.GetTableKey(t), id.String(), entity.ToJSONB())
    } else {
        impl.Cache.Set(impl.GetEntityKey(t, id, simple), entity.ToJSONB(), ENTITY_CACHE_EXPIRATION)
    }
}

func (impl *BaseCacheServiceImpl) MultiCacheEntity(t DictTypeEnum, entities []EntityIfc, simple bool) {
    if entities == nil || len(entities) == 0 {
        return
    }
    var m = make(map[string]interface{})
    if t.UseHashCache() {
        for _, entity := range entities {
            m[entity.GetBid().String()] = entity.ToJSONB()
        }
        impl.Cache.HMSet(impl.GetTableKey(t), m)
    } else {
        for _, entity := range entities {
            m[impl.GetEntityKey(t, entity.GetBid(), simple)] = entity.ToJSONB()
        }
        impl.Cache.MSet(m)
    }
}

func (impl *BaseCacheServiceImpl) CacheEmptyEntity(t DictTypeEnum, bid BID) {
    if t.UseHashCache() {
        impl.Cache.HSet(impl.GetTableKey(t), bid.String(), EMPTY_ENTITY_JSON_BYTES)
    } else {
        impl.Cache.Pipelined(func(pipe Pipeliner) error {
            pipe.Set(impl.GetEntityKey(t, bid, false), EMPTY_ENTITY_JSON_BYTES, ENTITY_CACHE_EXPIRATION)
            pipe.Set(impl.GetEntityKey(t, bid, true), EMPTY_ENTITY_JSON_BYTES, ENTITY_CACHE_EXPIRATION)
            return nil
        })
    }
}

func (impl *BaseCacheServiceImpl) DeleteEntityCache(t DictTypeEnum, bid BID, simple bool) {
    if t.UseHashCache() {
        impl.Cache.HDel(impl.GetTableKey(t), bid.String())
    } else {
        impl.Cache.Del(impl.GetEntityKey(t, bid, simple))
    }
}

func (impl *BaseCacheServiceImpl) MultiDeleteEntityCache(t DictTypeEnum, bids []BID, simple bool) {
    size := len(bids)
    keys := make([]string, size)
    if t.UseHashCache() {
        for i, bid := range bids {
            keys[i] = bid.String()
        }
        impl.Cache.HDel(impl.GetTableKey(t), keys...)
    } else {
        for i, bid := range bids {
            keys[i] = impl.GetEntityKey(t, bid, simple)
        }
        impl.Cache.Del(keys...)
    }
}

func (impl *BaseCacheServiceImpl) GetEntityCache(t DictTypeEnum, bid BID, simple bool) EntityIfc {
    var bytes []byte
    if t.UseHashCache() {
        bytes = impl.Cache.HGet(impl.GetTableKey(t), bid.String())
    } else {
        bytes = impl.Cache.Get(impl.GetEntityKey(t, bid, simple))
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
    return fmt.Sprintf("Dict:Entity:%s:%s:%v", t.String(), simpleStr, bid)
}

func (impl *BaseCacheServiceImpl) GetTableKey(t DictTypeEnum) string {
    return fmt.Sprintf("Dict:EntityHash:%s", t.String())
}

func (impl *BaseCacheServiceImpl) MultiGetEntityCache(t DictTypeEnum, bids []BID, simple bool) []EntityIfc {
    var size = len(bids)
    var jsons []interface{}
    if t.UseHashCache() {
        var bidStrs = make([]string, size)
        for i, bid := range bids {
            bidStrs[i] = bid.String()
        }
        jsons = impl.Cache.HMGet(impl.GetTableKey(t), bidStrs...)
    } else {
        var keys = make([]string, size)
        for i, bid := range bids {
            keys[i] = impl.GetEntityKey(t, bid, simple)
        }
        jsons = impl.Cache.MGet(keys...)
    }
    var entities = make([]EntityIfc, size)
    for i, str := range jsons {
        if str != nil {
            entities[i] = t.ParseJSON(str.(string))
        }
    }
    return entities
}
