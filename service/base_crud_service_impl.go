package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence/entity"
    "github.com/Danceiny/go.utils"
)

type BaseCrudServiceImpl struct {
    RepoServ       RepositoryService
    CacheServ      BaseCacheService
    IdFirewallServ IdFirewallService
}

func (impl *BaseCrudServiceImpl) Delete(entity EntityIfc) {

}
func (impl *BaseCrudServiceImpl) Update(entity EntityIfc) {

}
func (impl *BaseCrudServiceImpl) Add(entity EntityIfc) {
    impl.RepoServ.Add(entity)
}

func (impl *BaseCrudServiceImpl) MultiGet(t DictTypeEnum, bids []BID, simple bool) []EntityIfc {
    var entities = impl.CacheServ.MultiGetEntityCache(t, bids, simple)
    var entitiesNeedToCache = make([]EntityIfc, 0, 16)
    var i = 0
    for l := len(bids); i < l; i++ {
        if entities[i] == nil || go_utils.InterfaceHasNilValue(entities[i]) {
            var bid = bids[i]
            if impl.IdFirewallServ.IsBlackId(CATEGORY, bid) {
                continue
            }
            var entity = impl.RepoServ.Get(t, bid, simple, false)
            if entity == nil {
                // special case:
                // when query CATEGORY and miss cache and not in black list,
                // and db say not found,
                // we put the id to black list
                if !impl.IdFirewallServ.BlackingId(t, bid) {
                    // if we don't blacking the id, we cache a empty entity to identify that it's not existed
                    impl.CacheServ.CacheEmptyEntity(t, bids[i])
                }
            } else {
                entities[i] = entity
                entitiesNeedToCache = append(entitiesNeedToCache, entity)
            }
        } else if entities[i].IsEmpty() {
            // convert empty entity to null
            entities[i] = nil
        }
    }
    if len(entitiesNeedToCache) != 0 {
        impl.CacheServ.MultiCacheEntity(t, entitiesNeedToCache, simple)
    }
    return entities
}
func (impl *BaseCrudServiceImpl) Get(t DictTypeEnum, bid BID) EntityIfc {
    return nil
}
func (impl *BaseCrudServiceImpl) GetEntity(t DictTypeEnum, bid BID,
    simple bool, skipCache bool, withTrashed bool) EntityIfc {
    var entity EntityIfc
    entity = impl.CacheServ.GetEntityCache(t, bid, false)
    if entity == nil {
        entity = impl.RepoServ.Get(t, bid, false, false)
        if entity == nil {
            return nil
        } else {
            impl.CacheServ.CacheEntity(t, entity, false)
        }
    }
    return entity
}
func (impl *BaseCrudServiceImpl) Exist(t DictTypeEnum, bid BID) bool {
    return true
}
