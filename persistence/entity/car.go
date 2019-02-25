package entity

import (
    . "github.com/Danceiny/dict-service/common"
    "github.com/Danceiny/go.fastjson"
)

type CarLevel int

const (
    _ CarLevel = iota
    KIND
    BRAND
    SERIES
    MODEL
)

type CarEntity struct {
    TreeEntity
    Pid   BID `gorm:"column:parent_bid" json:"parentBid"`
}

func (CarEntity) GetType() DictTypeEnum {
    return CAR
}

func (entity *CarEntity) GetPid() BID {
    return entity.Pid
}

func (entity *CarEntity) GetDefaultBid() BID {
    // abstract method, implemented by subclass
    return INT(-1)
}

func (entity *CarEntity) ToJSONB() []byte {
    return fastjson.ToJSONB(entity)
}
