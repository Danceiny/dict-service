package service

import (
    "fmt"
    "github.com/Danceiny/dict-service/common/FastJson"
    . "github.com/Danceiny/dict-service/persistence"
    . "github.com/Danceiny/dict-service/persistence/entity"
    "time"
)

var (
    BaseCacheServiceImplCpt       *BaseCacheServiceImpl
    REDIS_ENTITY_CACHE_EXPIRATION time.Duration
)

func init() {
    BaseCacheServiceImplCpt = &BaseCacheServiceImpl{}
    REDIS_ENTITY_CACHE_EXPIRATION = 30 * 24 * time.Hour
}

type BaseCacheService interface {
    CacheEntity(t DictTypeEnum, entity Entity, simple bool)
    MultiCacheEntity(t DictTypeEnum, entities []Entity, simple bool)
    CacheEmptyEntity(t DictTypeEnum, bid BID);
    DeleteEntityCache(t DictTypeEnum, bid BID, simple bool);
    MultiDeleteEntityCache(t DictTypeEnum, bids []BID, simple bool)
    GetEntityCache(t DictTypeEnum, bid BID, simple bool) Entity
    GetEntityKey(t DictTypeEnum, bid BID, simple bool) string
    GetTableKey(t DictTypeEnum) string
}

type BaseCacheServiceImpl struct {
}

func (impl *BaseCacheServiceImpl) CacheEntity(t DictTypeEnum, entity Entity, simple bool) {
    var id = entity.GetBid()
    if t.UseHashCache() {
        RedisImplCpt.HSet(impl.GetTableKey(t), fmt.Sprint(id), FastJson.ToJSON(entity))
    } else {
        RedisImplCpt.Set(impl.GetEntityKey(t, id, simple), FastJson.ToJSON(entity), REDIS_ENTITY_CACHE_EXPIRATION)
    }

}
func (impl *BaseCacheServiceImpl) MultiCacheEntity(entities []Entity, simple bool) {

}
func (impl *BaseCacheServiceImpl) CacheEmptyEntity(t DictTypeEnum, bid BID) {

}
func (impl *BaseCacheServiceImpl) DeleteEntityCache(t DictTypeEnum, bid BID, simple bool) {

}
func (impl *BaseCacheServiceImpl) MultiDeleteEntityCache(t DictTypeEnum, bids []BID, simple bool) {

}
func (impl *BaseCacheServiceImpl) GetEntityCache(t DictTypeEnum, bid BID, simple bool) Entity {
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
    return fmt.Sprintf("Dict:Entity:%s:%s:%v", t, simpleStr, bid)
}

func (impl *BaseCacheServiceImpl) GetTableKey(t DictTypeEnum) string {
    return fmt.Sprintf("Dict:EntityHash:%s", t.String())
}
