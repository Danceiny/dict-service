package entity

import (
    . "github.com/Danceiny/dict-service/common"
    "github.com/Danceiny/go.fastjson"
)

type CategoryLevel int

const (
    _ CategoryLevel = iota
    ROOT
    FIRST
    SECOND
    THIRD
)

type CategoryEntity struct {
    TreeEntity
    Bid    STRING `gorm:"column:bid" json:"bid"`
    Pid    STRING `gorm:"column:parent_bid" json:"parentBid"`
    Pinyin string
}

func (entity CategoryEntity) GetBid() BID {
    return entity.Bid
}

func (CategoryEntity) GetType() DictTypeEnum {
    return CATEGORY
}

func (entity *CategoryEntity) GetPid() BID {
    return entity.Pid
}

func (entity *CategoryEntity) GetDefaultBid() BID {
    return STRING("")
}

func (entity *CategoryEntity) ToJSONB() []byte {
    return fastjson.ToJSONB(entity)
}
