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
