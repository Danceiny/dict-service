package service

import (
	"fmt"
	. "github.com/Danceiny/dict-service/persistence"
	. "github.com/Danceiny/dict-service/persistence/entity"
)

const (
	PARENT   = 0
	CHILDREN = 1
)

var (
	TreeCacheServiceImplCpt *TreeCacheServiceImpl
)

func init() {
	TreeCacheServiceImplCpt = &TreeCacheServiceImpl{
		RedisImplCpt, BaseCacheServiceImplCpt}
}

type TreeCacheServiceImpl struct {
	redis            Redis
	baseCacheService BaseCacheService
}

func (impl *TreeCacheServiceImpl) GetParentBids(t DictTypeEnum, bid BID) []BID {
	return nil
}
func (impl *TreeCacheServiceImpl) GetParentBranchKey(t DictTypeEnum) string {
	return impl.getBranchKey(t, PARENT)
}
func (impl *TreeCacheServiceImpl) MultiGetParentBids(t DictTypeEnum, bids []BID) [][]BID {
	return nil
}
func (impl *TreeCacheServiceImpl) SetParentBids(t DictTypeEnum, bid BID, pids []BID) {

}
func (impl *TreeCacheServiceImpl) DeleteParentBids(t DictTypeEnum, bid BID) {

}
func (impl *TreeCacheServiceImpl) MultiDeleteParentBids(t DictTypeEnum, bids []BID) {

}
func (impl *TreeCacheServiceImpl) GetChildrenBranchKey(t DictTypeEnum) string {
	return impl.getBranchKey(t, CHILDREN)
}
func (impl *TreeCacheServiceImpl) GetChildrenBids(t DictTypeEnum, bid BID) {

}
func (impl *TreeCacheServiceImpl) MultiGetChildrenBids(t DictTypeEnum, bids []BID) [][]BID {
	return nil
}
func (impl *TreeCacheServiceImpl) SetChildrenBids(t DictTypeEnum, bid BID, cids []BID) {

}
func (impl *TreeCacheServiceImpl) DeleteChildrenBids(t DictTypeEnum, bid BID) {

}
func (impl *TreeCacheServiceImpl) GetEntityInPipeline(t DictTypeEnum, bid BID, simple bool, p int, c int) [][]byte {
	strings := impl.redis.Pipelined(func(pipe Pipeliner) error {
		if t.UseHashCache() {
			pipe.HGet(impl.baseCacheService.GetTableKey(t), bid.String())
		} else {
			pipe.Get(impl.baseCacheService.GetEntityKey(t, bid, simple))
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
