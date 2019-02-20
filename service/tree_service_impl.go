package service

import (
	"fmt"
	. "github.com/Danceiny/dict-service/api"
	. "github.com/Danceiny/dict-service/common"
	. "github.com/Danceiny/dict-service/persistence"
	. "github.com/Danceiny/dict-service/persistence/entity"
	log "github.com/sirupsen/logrus"
)

var (
	MAX_DEPTH = 1024
)

type TreeServiceImpl struct {
	RepoServ      RepositoryService
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
		if size > 1 && bytesList[1] != nil {
			entity.SetPids(t.ParseBids(bytesList[1]))
		}
		impl.LoadParent(entity, p, simple)
	}
	if c != 0 {
		if size > 1 && bytesList[size-1] != nil {
			entity.SetCids(t.ParseBids(bytesList[size-1]))
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
		for loopCounter--; bidZero != pid && loopCounter > 0; {
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
	onlyCache bool) []TreeEntityIfc {
	// todo
	return nil
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
