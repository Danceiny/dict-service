package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence/entity"
)

type BaseCrudService interface {
    Delete(entity EntityIfc)
    Update(entity EntityIfc)
    Add(entity EntityIfc)
    MultiGet(t DictTypeEnum, bids []BID, simple bool) []EntityIfc
    Get(t DictTypeEnum, bid BID) EntityIfc
    GetEntity(t DictTypeEnum, bid BID,
        simple bool, skipCache bool, withTrashed bool) EntityIfc
    Exist(t DictTypeEnum, bid BID) bool
}
