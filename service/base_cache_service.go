package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence/entity"
)

type BaseCacheService interface {
    CacheEntity(t DictTypeEnum, entity EntityIfc, simple bool)
    MultiCacheEntity(t DictTypeEnum, entities []EntityIfc, simple bool)
    MultiGetEntityCache(t DictTypeEnum, bid []BID, simple bool) []EntityIfc
    CacheEmptyEntity(t DictTypeEnum, bid BID)
    DeleteEntityCache(t DictTypeEnum, bid BID, simple bool)
    MultiDeleteEntityCache(t DictTypeEnum, bids []BID, simple bool)
    GetEntityCache(t DictTypeEnum, bid BID, simple bool) EntityIfc
    GetEntityKey(t DictTypeEnum, bid BID, simple bool) string
    GetTableKey(t DictTypeEnum) string
}
