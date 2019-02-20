package entity

import (
	. "github.com/Danceiny/dict-service/common"
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
	Name  string
	Level NodeLevel
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
