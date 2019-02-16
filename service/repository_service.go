package service

import (
	. "github.com/Danceiny/dict-service/persistence/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type RepositoryService interface {
	Add(ifc EntityIfc)
	Get(t DictTypeEnum, bid BID, simple bool, withTrashed bool) EntityIfc
	GetEntity(t DictTypeEnum, bid BID, withTrashed bool) EntityIfc
	GetCids(t DictTypeEnum, parentBid BID) []BID
}
