package entity

import (
    . "github.com/Danceiny/dict-service/common"
)

type TreeEntityIfc interface {
    EntityIfc
    GetPid() BID
    GetLevel() int
    SetLevel(level int)
    GetCids() []BID
    SetCids(bids []BID)
    GetPids() []BID
    SetPids(bids []BID)
    GetParent() TreeEntityIfc
    SetParent(parent TreeEntityIfc)

    GetParentChain() []TreeEntityIfc
    SetParentChain(parentChain []TreeEntityIfc)
    GetChildren() []TreeEntityIfc
    SetChildren(children []TreeEntityIfc)

    GetDefaultBid() BID

    GetOldLevel() *int
    GetOldPid() BID
}

type TreeEntity struct {
    BaseEntity
    Level       int             `gorm:"column:node_level" json:"level"`
    Name        string          `gorm:"column:node_name" json:"name"`
    Cids        []BID           `gorm:"-" json:"-"`
    Pids        []BID           `gorm:"-" json:"-"`
    Children    []TreeEntityIfc `gorm:"-" json:"-"`
    Parent      TreeEntityIfc   `gorm:"-" json:"-"`
    ParentChain []TreeEntityIfc `gorm:"-" json:"-"`
    OldLevel    *int            `gorm:"-" json:"-"`
    OldPid      BID             `gorm:"-" json:"-"`
    Weight      int             `gorm:"column:weight" json:"weight"`
    // for cache fields (java fastjson/jackson)
    Bid       interface{} `gorm:"-" json:"bid"`
    LevelName string      `gorm:"-" json:"levelEnum"`
}

func (entity *TreeEntity) GetType() DictTypeEnum {
    // abstract method, implemented by subclass
    return COMMON
}

func (entity *TreeEntity) GetPid() BID {
    // abstract method, implemented by subclass
    return nil
}

func (entity *TreeEntity) GetDefaultBid() BID {
    // abstract method, implemented by subclass
    return nil
}

func (entity *TreeEntity) GetOldPid() BID {
    return entity.OldPid
}

func (entity *TreeEntity) GetOldLevel() *int {
    return entity.OldLevel
}

func (entity *TreeEntity) GetLevel() int {
    return entity.Level
}

func (entity *TreeEntity) SetLevel(level int) {
    entity.Level = level
}

func (entity *TreeEntity) GetParent() TreeEntityIfc {
    return entity.Parent
}
func (entity *TreeEntity) SetParent(parent TreeEntityIfc) {
    entity.Parent = parent
}
func (entity *TreeEntity) GetParentChain() []TreeEntityIfc {
    return entity.ParentChain
}
func (entity *TreeEntity) SetParentChain(parentChain []TreeEntityIfc) {
    entity.ParentChain = parentChain
}
func (entity *TreeEntity) GetChildren() []TreeEntityIfc {
    return entity.Children
}
func (entity *TreeEntity) SetChildren(children []TreeEntityIfc) {
    entity.Children = children
}
func (entity *TreeEntity) GetCids() []BID {
    return entity.Cids
}
func (entity *TreeEntity) SetCids(bids []BID) {
    entity.Cids = bids
}
func (entity *TreeEntity) GetPids() []BID {
    return entity.Pids
}
func (entity *TreeEntity) SetPids(bids []BID) {
    entity.Pids = bids
}
