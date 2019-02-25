package service

import (
    "fmt"
    . "github.com/Danceiny/dict-service/api"
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence"
    . "github.com/Danceiny/dict-service/persistence/entity"
    utils "github.com/Danceiny/go.utils"
    log "github.com/sirupsen/logrus"
)

var (
    MAX_DEPTH = 1024
)

type TreeServiceImpl struct {
    RepoServ      RepositoryService
    BaseCacheServ BaseCacheService
    BaseCrudServ  BaseCrudService
    TreeCacheServ TreeCacheService
}

func (impl *TreeServiceImpl) GetTree(t DictTypeEnum, bid BID, p int, c int,
    simple bool, skipCache bool) TreeEntityIfc {
    log.Infof("GetTree t: %s, bid: %v, p:%d, c:%d", t.String(), bid, p, c)
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
        log.Println("loadParent")
        if size > 1 && bytesList[1] != nil {
            entity.SetPids(t.ParseBidsB(bytesList[1]))
        }
        impl.LoadParent(entity, p, simple)
    }
    if c != 0 {
        log.Println("loadChildren")
        if size > 1 && bytesList[size-1] != nil {
            entity.SetCids(t.ParseBidsB(bytesList[size-1]))
        } else {
            var cids = impl.RepoServ.GetCids(t, bid)
            entity.SetCids(cids)
            impl.TreeCacheServ.SetChildrenBids(t, bid, cids)
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
    var oldLevel = entity.GetOldLevel()
    if oldLevel != nil && *oldLevel != entity.GetLevel() {
        // 注：升降级level后，不处理children
        pid := entity.GetPid()
        parent := impl.getEntity(t, pid,
            false, true, false)
        if parent == nil {
            ThrowArgException(fmt.Sprintf("new parent not found by id: %v", pid))
        }
        if entity.GetLevel() < parent.GetLevel() {
            ThrowArgException("level higher than parent.level is not allowed")
        }
        impl.LoadChildren(entity, 1, true)
        for _, child := range entity.GetChildren() {
            if entity.GetLevel() > child.GetLevel() {
                ThrowArgException("level lower than child.level is not allowed")
            }
        }
    }
    // 修改了pid, 警告：如果修改的是树顶部的一些节点，则下面的
    if entity.GetOldPid() != nil && entity.GetOldPid() != entity.GetPid() {
        impl.changeNodeParent(t, entity.GetBid(), entity.GetPid(), entity.GetOldPid())
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
    if depth == 0 {
        return
    }
    if depth < 0 {
        depth = MAX_DEPTH
    }
    var pid = entity.GetPid()
    var bidZero = entity.GetDefaultBid()
    if bidZero == pid {
        entity.SetPids(make([]BID, 0))
        entity.SetParentChain(make([]TreeEntityIfc, 0))
        return
    }
    var level = entity.GetLevel()
    var bid = entity.GetBid()
    var t = entity.GetType()
    var nodes = make([]TreeEntityIfc, 0, 3)
    var pids = entity.GetPids()
    // 从缓存中取pids
    if pids == nil {
        pids = impl.TreeCacheServ.GetParentBids(t, bid)
        entity.SetPids(pids)
    }
    if cap(pids) == 0 {
        // 没有pids的缓存，先从数据库build
        pids = make([]BID, 0, 3)
        var parent TreeEntityIfc
        // 为了防止特殊情况下脏数据影响递归
        var loopCounter = MAX_DEPTH
        for loopCounter--; !bidZero.Equal(pid) && loopCounter > 0; {
            var parentVar1 = impl.BaseCrudServ.GetEntity(t, pid, false, simple, false)
            if parentVar1 != nil {
                parent = parentVar1.(TreeEntityIfc)
                pid = parent.GetPid()
                // for cache
                pids = append(pids, parent.GetBid())
                nodes = append(nodes, parent)
                if pid == parent.GetBid() {
                    // 如果parent的pid指向了parent自己，那就到头了
                    break
                }
                if level < parent.GetLevel() {
                    // 应该是数据异常，理论上该节点的level值要不小于parent的level
                    log.Error("【警告】数据层级异常，请排查：{}-{}， level({}), parent: {}",
                        t, bid, level, parent)
                    break
                }
                level = parent.GetLevel()
            } else {
                log.Warning("find parents by bid not found: {}", pid)
                break
            }
        }
        if loopCounter <= 0 {
            log.Error("【警告】递归深度异常，请排查：{}-{}", t, bid)
        }
        // 缓存整个pids
        entity.SetPids(pids)
        impl.TreeCacheServ.SetParentBids(t, bid, pids)
        // 返回部分parent
        var toIndex = len(pids)
        if depth < toIndex {
            toIndex = depth
        }
        entity.SetParentChain(nodes[:toIndex])
        return
    }
    entity.SetPids(pids)
    if len(pids) == 0 {
        entity.SetParentChain(make([]TreeEntityIfc, 0))
        return
    }
    var toIndex = len(pids)
    if depth < toIndex {
        toIndex = depth
    }
    var list = impl.multiGetBaseTreeEntities(t, (pids)[:toIndex], simple)
    for _, p := range list {
        if p == nil {
            impl.TreeCacheServ.DeleteParentBids(t, bid)
        } else {
            nodes = append(nodes, p)
        }
    }
    entity.SetParentChain(nodes)
}

func (impl *TreeServiceImpl) LoadChildren(entity TreeEntityIfc, depth int, simple bool) {
    if depth == 0 {
        return
    }
    if depth < 0 {
        depth = MAX_DEPTH
    }
    var bid = entity.GetBid()
    var t = entity.GetType()
    // 此处为减少redis请求数，使用multi-get将cids绑到entity身上
    var cids = entity.GetCids()
    // 从缓存中取cids
    if cap(cids) == 0 {
        cids = impl.TreeCacheServ.GetChildrenBids(t, bid)
        entity.SetCids(cids)
    }
    depth--
    // cids无缓存，读数据库
    if cap(cids) == 0 {
        cids = make([]BID, 0, 16)
        var children = impl.getChildrenFromDB(t, bid, simple)
        for _, child := range children {
            cids = append(cids, child.GetBid())
            impl.LoadChildren(child, depth, simple)
        }
        entity.SetChildren(children)
        entity.SetCids(cids)
    } else if len(cids) == 0 {
        entity.SetChildren(make([]TreeEntityIfc, 0))
    } else {
        // 一次获取所有的children
        var var1 = impl.multiGetBaseTreeEntities(t, cids, simple);
        // children2不含null
        var children2 = make([]TreeEntityIfc, 0, 16)
        for _, var2 := range var1 {
            if var2 == nil { // InterfaceHasNilValue ?
                // 缓存的cids，对应的实体拿不到，说明有数据异常
                impl.TreeCacheServ.DeleteChildrenBids(t, bid);
            } else {
                children2 = append(children2, var2)
            }
        }
        var l = len(children2)
        if depth > 0 && l > 0 {
            // 如果深度>=2（对应depth>0条件），则获取所有children的cids（cidsList）
            var bids = make([]BID, 0, 16)
            for _, child := range children2 {
                bids = append(bids, child.GetBid())
            }
            var cidsList = impl.GetCids(t, bids);
            // 将cidsList打平，一次获取全部
            var indexes = make([][2]int, l)
            // mixIds：所有孙子的id数组
            var mixIds = make([]BID, 0, 16)
            var i = 0
            for c := 0; i < l; i++ {
                var e = children2[i]
                var ids = cidsList[i];
                mixIds = append(mixIds, ids...)
                e.SetCids(ids)
                indexes[i][0] = c
                c += len(ids)
                indexes[i][1] = c
            }
            // 深度-1，准备给所有的children获取children
            depth--
            // 一次获取所有的“孙子”（顺序与mixIds一一对应）
            var entityList = impl.multiGetBaseTreeEntities(t, mixIds, simple);
            for i := 0; i < l; i++ {
                var child = children2[i]
                var from = indexes[i][0]
                var to = indexes[i][1]
                child.SetChildren(entityList[from:to])
            }
            // 下面的还可以继续优化
            if depth > 0 {
                for _, child2 := range children2 {
                    if child2 != nil && !utils.InterfaceHasNilValue(child2) {
                        var children3 = child2.GetChildren()
                        for _, child3 := range children3 {
                            if child3 != nil && !utils.InterfaceHasNilValue(child3) {
                                impl.LoadChildren(child3, depth, simple)
                            } else {
                                log.Errorf("nil %v", child3)
                            }
                        }
                    } else {
                        log.Errorf("nil %v", child2)
                    }
                }
            }
        }
        entity.SetChildren(children2)
    }
}

func (impl *TreeServiceImpl) getChildrenFromDB(t DictTypeEnum, bid BID, simple bool) []TreeEntityIfc {
    var entities = impl.RepoServ.GetByPid(t, bid, simple)
    impl.BaseCacheServ.MultiCacheEntity(t, entities, simple)
    var ret = make([]TreeEntityIfc, len(entities))
    for i, entity := range entities {
        ret[i] = entity.(TreeEntityIfc)
    }
    return ret
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
                if entity.GetDefaultBid() != entity.GetPid() {
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
    onlyCache bool, skipCache bool) []TreeEntityIfc {
    var size = len(bids)
    var entities []TreeEntityIfc
    if onlyCache {
        var var1 = impl.BaseCacheServ.MultiGetEntityCache(t, bids, simple)
        entities = make([]TreeEntityIfc, len(var1))
        for i, var2 := range var1 {
            if var2 != nil && var2.IsEmpty() {
                entities[i] = var2.(TreeEntityIfc)
            }
        }
    } else {
        entities = impl.multiGetBaseTreeEntities(t, bids, simple)
    }
    if p == 0 && c == 0 {
        return entities
    }
    var pidsList [][]BID
    var cidsList [][]BID
    if p != 0 {
        pidsList = make([][]BID, size)
        var var1 = impl.GetPids(t, bids)
        for i, var2 := range var1 {
            if cap(var2) == 0 {
                continue
            }
            if p < 0 {
                pidsList[i] = var2
                continue
            }
            var toIndexMax = len(var2)
            var toIndex = p
            if toIndex > toIndexMax {
                toIndex = toIndexMax
            }
            pidsList[i] = var2[0:toIndex]
        }
    }
    if c != 0 {
        cidsList = impl.GetCids(t, bids)
    }
    for i := 0; i < size; i++ {
        var entity = entities[i]
        if entity == nil {
            continue
        }
        if p != 0 {
            entity.SetPids(pidsList[i])
        }
        if c != 0 {
            entity.SetCids(cidsList[i])
        }
    }
    if !onlyId {
        // 为了一次获取全部的parent和children，需要indexes来索引这些parent和children的位置
        // indexes: [[parentStartIndex, parentEndIndex, childrenStartIndex, childrenEndIndex], ...]
        // 与entities、bids一一保持一一对应关系
        var indexes = make([][4]int, size)
        var flatIdsList = make([]BID, 0, 96)
        var cursor = 0
        var pidsListIsNotEmpty = len(pidsList) != 0
        var cidsListIsNotEmpty = len(cidsList) != 0
        for i := 0; i < size; i++ {
            if pidsListIsNotEmpty {
                var var1 = pidsList[i]
                if cap(var1) == 0 {
                    continue
                }
                indexes[i][0] = cursor
                cursor += len(var1)
                indexes[i][1] = cursor
                flatIdsList = append(flatIdsList, var1...)
            }
            if cidsListIsNotEmpty {
                var var2 = cidsList[i]
                if cap(var2) == 0 {
                    continue
                }
                indexes[i][2] = cursor
                cursor += len(var2)
                indexes[i][3] = cursor
                flatIdsList = append(flatIdsList, var2...)
            }
        }
        var parentsAndChildren = impl.multiGetBaseTreeEntities(t, flatIdsList, simple)
        for i := 0; i < size; i++ {
            var entity = entities[i]
            if entity == nil {
                continue
            }
            if p != 0 {
                entity.SetParentChain(
                    parentsAndChildren[indexes[i][0]:indexes[i][1]])
            }
            if c != 0 {
                entity.SetChildren(parentsAndChildren[indexes[i][2]:indexes[i][3]])
            }
        }
    } else {
        for i := 0; i < size; i++ {
            var entity = entities[i]
            if entity != nil {
                entity.SetParentChain(nil)
                entity.SetChildren(nil)
            }
        }
    }
    return entities
}

func (impl *TreeServiceImpl) multiGetBaseTreeEntities(t DictTypeEnum, bids []BID, simple bool) []TreeEntityIfc {
    var entities = impl.BaseCrudServ.MultiGet(t, bids, simple)
    var ret = make([]TreeEntityIfc, len(entities))
    for i, entity := range entities {
        if entity != nil {
            ret[i] = entity.(TreeEntityIfc)
        }
    }
    return ret
}
