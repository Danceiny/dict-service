package service

import (
    "fmt"
    . "github.com/Danceiny/dict-service/api"
    . "github.com/Danceiny/dict-service/common"
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
        tmp := impl.BaseCrudServ.GetEntity(t, bid, simple, skipCache, false)
        if tmp == nil {
            return nil
        } else {
            return tmp.(TreeEntityIfc)
        }
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
        impl.LoadParent(entity, p, simple)
    }
    if c != 0 {
        if size > 1 && bytesList[size-1] != nil {
            entity.SetCids(ParseBids(bytesList[size-1]))
        } else {
            entity.SetCids(impl.RepoServ.GetCids(t, bid))
            impl.TreeCacheServ.SetChildrenBids(t, bid, entity.GetCids())
        }
        impl.LoadChildren(entity, c, simple)
    }
    return entity
}

func (impl *TreeServiceImpl) getEntity(t DictTypeEnum, bid BID,
    simple bool, skipCache bool, withTrashed bool) TreeEntityIfc {
    var1 := impl.BaseCrudServ.GetEntity(t, bid,
        false, true, false)
    if var1 == nil {
        return nil
    }
    return var1.(TreeEntityIfc)
}

func (impl *TreeServiceImpl) Save(entity TreeEntityIfc) {
    t := entity.GetType()
    // todo: check id (not bid)4
    // 修改了level，这是一件大事
    if entity.GetOldLevel() != nil && entity.GetOldLevel() != entity.GetLevel() {
        // 注：升降级level后，不处理children
        pid := entity.GetParentBid()
        parent := impl.getEntity(t, pid,
            false, true, false)
        if parent == nil {
            ThrowArgException(fmt.Sprintf("new parent not found by id: %v", pid))
        }
        if NodeLevelAltB(entity.GetLevel(), parent.GetLevel()) {
            ThrowArgException("level higher than parent.level is not allowed")
        }
        impl.LoadChildren(entity, 1, true)
        for _, child := range entity.GetChildren() {
            if NodeLevelAgtB(entity.GetLevel(), child.GetLevel()) {
                ThrowArgException("level lower than child.level is not allowed")
            }
        }
    }
    // 修改了pid, 警告：如果修改的是树顶部的一些节点，则下面的
    if entity.GetOldParentBid() != nil && entity.GetOldParentBid() != entity.GetParentBid() {
        impl.changeNodeParent(t, entity.GetBid(), entity.GetParentBid(), entity.GetOldParentBid())
    }
    impl.BaseCrudServ.Update(entity)
}

// todo
func (impl *TreeServiceImpl) changeNodeParent(t DictTypeEnum, bid BID, npid BID, opid BID) {

}

func (impl *TreeServiceImpl) Add(entity TreeEntityIfc) {

}
func (impl *TreeServiceImpl) UpdateCommonProps(entity TreeEntityIfc, req *TreeUpdateReq) {

}
func (impl *TreeServiceImpl) AdjustSortedWeight(sortedEntities *[]TreeEntityIfc) {

}
func (impl *TreeServiceImpl) Delete(t DictTypeEnum, bid BID) {

}

func (impl *TreeServiceImpl) LoadParent(entity TreeEntityIfc, depth int, simple bool) {

}

func (impl *TreeServiceImpl) LoadChildren(entity TreeEntityIfc, depth int, simple bool) {
}

func (impl *TreeServiceImpl) GetCids(t DictTypeEnum, bids []BID) [][]BID {
    ret := impl.TreeCacheServ.MultiGetChildrenBids(t, bids)
    for i, var1 := range ret {
        if var1 == nil {
            var2 := impl.RepoServ.GetCids(t, bids[i])
            ret[i] = var2
            impl.TreeCacheServ.SetChildrenBids(t, bids[i], var2)
        }
    }
    return ret
}
func (impl *TreeServiceImpl) GetPids(t DictTypeEnum, bids []BID) [][]BID {
    ret := impl.TreeCacheServ.MultiGetParentBids(t, bids)
    for i, var1 := range ret {
        if var1 == nil {
            bid := bids[i]
            var2 := impl.BaseCrudServ.GetEntity(t, bid, true, false, false)
            if var2 != nil {
                entity := var2.(TreeEntityIfc)
                if entity.GetDefaultBid() != entity.GetParentBid() {
                    impl.LoadParent(entity, -1, true)
                    pids := entity.GetPids()
                    ret[i] = pids
                    impl.TreeCacheServ.SetParentBids(t, bid, pids)
                } else {
                    ret[i] = make([]BID, 0)
                }
            }
        }
    }
    return ret
}
func (impl *TreeServiceImpl) MultiGet(t DictTypeEnum, bids []BID,
    simple bool,
    p int, c int,
    onlyId bool,
    onlyCache bool) []TreeEntityIfc {
    // todo
    return nil
}
