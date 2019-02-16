package service

import (
    . "github.com/Danceiny/dict-service/api"
    . "github.com/Danceiny/dict-service/persistence"
    . "github.com/Danceiny/dict-service/persistence/entity"
)

type TreeServiceImpl struct {
    RepoServ      RepositoryService
    BaseCrudServ  BaseCrudService
    TreeCacheServ TreeCacheService
}

func (impl *TreeServiceImpl) GetTree(t DictTypeEnum, bid BID, p int, c int,
    simple bool, skipCache bool) TreeEntityIfc {
    var perfCheck = p != 0 || c != 0
    if !perfCheck {
        return impl.BaseCrudServ.GetEntity(t, bid, simple, skipCache, false).(TreeEntityIfc)
    }
    var bytesList [][]byte
    bytesList = impl.TreeCacheServ.GetEntityInPipeline(t, bid, simple, p, c)
    var entity TreeEntityIfc
    if bytesList[0] == nil {
        entity = nil
    } else {
        entity = ParseEntityFromJSON(t, bytesList[0]).(TreeEntityIfc)
    }
    if entity == nil {
        // 缓存未命中
        tmp := impl.BaseCrudServ.GetEntity(t, bid, simple, true, false)
        if tmp == nil {
            return nil
        } else {
            entity = tmp.(TreeEntityIfc)
        }
    } else if entity.IsEmpty() {
        return nil
    }
    size := len(bytesList)
    if p != 0 {
        if size > 1 && bytesList[1] != nil {
            entity.SetPids(ParseBids(bytesList[1]))
        }
        impl.LoadParent(&entity, p, simple)
    }
    if c != 0 {
        if size > 1 && bytesList[size-1] != nil {
            entity.SetCids(ParseBids(bytesList[size-1]))
        } else {
            entity.SetCids(impl.RepoServ.GetCids(t, bid))
            impl.TreeCacheServ.SetChildrenBids(t, bid, entity.GetCids())
        }
        impl.LoadChildren(&entity, c, simple)
    }
    return entity
}
func (impl *TreeServiceImpl) Save(entity *TreeEntityIfc) {

}
func (impl *TreeServiceImpl) Add(entity *TreeEntityIfc) {

}
func (impl *TreeServiceImpl) UpdateCommonProps(entity *TreeEntityIfc, req *TreeUpdateReq) {

}
func (impl *TreeServiceImpl) AdjustSortedWeight(sortedEntities *[]TreeEntityIfc) {

}
func (impl *TreeServiceImpl) Delete(t DictTypeEnum, bid BID) {

}
func (impl *TreeServiceImpl) LoadParent(entity *TreeEntityIfc, depth int, simple bool) {

}
func (impl *TreeServiceImpl) LoadChildren(entity *TreeEntityIfc, depth int, simple bool) {
}
func (impl *TreeServiceImpl) GetCids(t DictTypeEnum, bids []BID) [][]BID {
    // todo
    return nil
}
func (impl *TreeServiceImpl) GetPids(t DictTypeEnum, bids []BID) [][]BID {
    // todo
    return nil
}
func (impl *TreeServiceImpl) MultiGet(t DictTypeEnum, bids []BID,
    simple bool,
    p int, c int,
    onlyId bool,
    onlyCache bool) []TreeEntityIfc {
    // todo
    return nil
}
