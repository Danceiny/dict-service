package service

import (
	. "github.com/Danceiny/dict-service/api"
	. "github.com/Danceiny/dict-service/persistence/entity"
)

type TreeService interface {
	GetTree(t DictTypeEnum, bid BID, p int, c int, simple bool, skipCache bool) TreeEntityIfc
	Save(entity *TreeEntityIfc)
	Add(entity *TreeEntityIfc)
	UpdateCommonProps(entity *TreeEntityIfc, req *TreeUpdateReq)
	AdjustSortedWeight(sortedEntities *[]TreeEntityIfc)
	Delete(t DictTypeEnum, bid BID)
	LoadParent(entity *TreeEntityIfc, depth int, simple bool)
	LoadChildren(entity *TreeEntityIfc, depth int, simple bool)
	GetCids(t DictTypeEnum, bids []BID) [][]BID
	GetPids(t DictTypeEnum, bids []BID) [][]BID
	MultiGet(t DictTypeEnum, bids []BID,
		simple bool,
		p int, c int,
		onlyId bool,
		onlyCache bool) []TreeEntityIfc
}
