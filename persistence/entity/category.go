package entity

import (
    . "github.com/Danceiny/dict-service/common"
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
    Bid         String
    Pid         String `gorm:"column:parent_bid" json:"parentBid"`
    Name        string
    Level       NodeLevel
    EnglishName string
    Pinyin      string
}

func (entity CategoryEntity) GetBid() BID {
    return entity.Bid
}

func (CategoryEntity) GetType() DictTypeEnum {
    return CATEGORY
}

func (entity *CategoryEntity) GetParentBid() BID {
    return entity.Pid
}

func (entity *CategoryEntity) GetDefaultBid() BID {
    return String("")
}
