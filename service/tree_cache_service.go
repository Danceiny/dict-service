package service

import (
	. "github.com/Danceiny/dict-service/persistence/entity"
)

type TreeCacheService interface {
	GetParentBids(t DictTypeEnum, bid BID) []BID
	GetParentBranchKey(t DictTypeEnum) string
	MultiGetParentBids(t DictTypeEnum, bids []BID) [][]BID
	SetParentBids(t DictTypeEnum, bid BID, pids []BID)
	DeleteParentBids(t DictTypeEnum, bid BID)
	MultiDeleteParentBids(t DictTypeEnum, bids []BID)
	GetChildrenBranchKey(t DictTypeEnum) string
	GetChildrenBids(t DictTypeEnum, bid BID)
	MultiGetChildrenBids(t DictTypeEnum, bids []BID) [][]BID
	SetChildrenBids(t DictTypeEnum, bid BID, cids []BID)
	DeleteChildrenBids(t DictTypeEnum, bid BID)
	GetEntityInPipeline(final DictTypeEnum, bid BID, simple bool, p int, c int) [][]byte
}
