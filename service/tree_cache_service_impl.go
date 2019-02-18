package service

import (
    "fmt"
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence/entity"
)

const (
    PARENT   = 0
    CHILDREN = 1
)

type TreeCacheServiceImpl struct {
    Cache         Redis
    BaseCacheServ BaseCacheService
}

func (impl *TreeCacheServiceImpl) GetParentBids(t DictTypeEnum, bid BID) []BID {
    c := impl.Cache.HGet(impl.GetChildrenBranchKey(t), bid.String())
    return ParseBids(c)
}

func (impl *TreeCacheServiceImpl) GetParentBranchKey(t DictTypeEnum) string {
    return impl.getBranchKey(t, PARENT)
}

// if return nil, caller still can expand the return value like `impl.bids2StrArr(bids)...`
func (impl *TreeCacheServiceImpl) bids2StrArr(bids []BID) []string {
    if bids == nil {
        return nil
    }
    bidStrs := make([]string, len(bids))
    for i, bid := range bids {
        bidStrs[i] = bid.String()
    }
    return bidStrs
}

func (impl *TreeCacheServiceImpl) MultiGetParentBids(t DictTypeEnum, bids []BID) [][]BID {
    c := impl.Cache.HMGet(impl.GetParentBranchKey(t), impl.bids2StrArr(bids)...)
    return MultiParseBids(c)
}

func (impl *TreeCacheServiceImpl) SetParentBids(t DictTypeEnum, bid BID, pids []BID) {
    impl.Cache.HSet(impl.GetParentBranchKey(t), bid.String(), Bids2Json(pids))
}

func (impl *TreeCacheServiceImpl) DeleteParentBids(t DictTypeEnum, bid BID) {
    impl.Cache.HDel(impl.GetParentBranchKey(t), bid.String())
}

func (impl *TreeCacheServiceImpl) MultiDeleteParentBids(t DictTypeEnum, bids []BID) {
    impl.Cache.HDel(impl.GetParentBranchKey(t), impl.bids2StrArr(bids)...)
}

func (impl *TreeCacheServiceImpl) GetChildrenBranchKey(t DictTypeEnum) string {
    return impl.getBranchKey(t, CHILDREN)
}

func (impl *TreeCacheServiceImpl) GetChildrenBids(t DictTypeEnum, bid BID) []BID {
    c := impl.Cache.HGet(impl.GetChildrenBranchKey(t), bid.String())
    return ParseBids(c)
}

func (impl *TreeCacheServiceImpl) MultiGetChildrenBids(t DictTypeEnum, bids []BID) [][]BID {
    c := impl.Cache.HMGet(impl.GetChildrenBranchKey(t), impl.bids2StrArr(bids)...)
    return MultiParseBids(c)
}

func (impl *TreeCacheServiceImpl) SetChildrenBids(t DictTypeEnum, bid BID, cids []BID) {
    impl.Cache.HSet(impl.GetChildrenBranchKey(t), bid.String(), Bids2Json(cids))
}

func (impl *TreeCacheServiceImpl) DeleteChildrenBids(t DictTypeEnum, bid BID) {
    impl.Cache.HDel(impl.GetChildrenBranchKey(t), bid.String())
}

func (impl *TreeCacheServiceImpl) GetEntityInPipeline(t DictTypeEnum, bid BID, simple bool, p int, c int) [][]byte {
    strings := impl.Cache.Pipelined(func(pipe Pipeliner) error {
        if t.UseHashCache() {
            pipe.HGet(impl.BaseCacheServ.GetTableKey(t), bid.String())
        } else {
            pipe.Get(impl.BaseCacheServ.GetEntityKey(t, bid, simple))
        }
        if p != 0 {
            pipe.HGet(impl.GetParentBranchKey(t), bid.String())
        }
        if c != 0 {
            pipe.HGet(impl.GetChildrenBranchKey(t), bid.String())
        }
        return nil
    })
    ret := make([][]byte, len(strings))
    for i, s := range strings {
        if s != "" {
            ret[i] = []byte(s)
        } else {
            ret[i] = nil
        }
    }
    return ret
}

func (impl *TreeCacheServiceImpl) getBranchKey(t DictTypeEnum, flag int) string {
    return fmt.Sprintf("Dict:%d:%s", flag, t.String())
}
