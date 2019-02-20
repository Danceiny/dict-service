package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence/entity"
)

type IdFirewallService interface {
    ValidateId(t DictTypeEnum, id BID) bool
    IsBlackId(t DictTypeEnum, id BID) bool
    BlackingId(t DictTypeEnum, id BID) bool
    UnblackingId(t DictTypeEnum, id BID) bool
    UnblackingDictType(t DictTypeEnum) bool
}
