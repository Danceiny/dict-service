package entity

import (
    . "github.com/Danceiny/dict-service/common"
)

type TreeEntityIfc interface {
    EntityIfc
    GetLevel() NodeLevel
    SetLevel(level NodeLevel)
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
}

type TreeEntity struct {
    BaseEntity
    Level       NodeLevel     `gorm:"column:node_level" json:"level"`
    Pid         BID           `gorm:"column:parent_bid" json:"parentBid"`
    Cids        []BID         `gorm:"-" json:"-"`
    Pids        []BID         `gorm:"-" json:"-"`
    Children    []*TreeEntity `gorm:"-" json:"-"`
    Parent      *TreeEntity   `gorm:"-" json:"-"`
    ParentChain []*TreeEntity `gorm:"-" json:"-"`
}

func (entity *TreeEntity) GetType() DictTypeEnum {
    // abstract method, implemented by subclass
    return 0
}

func (entity *TreeEntity) GetLevel() NodeLevel {
    return entity.Level
}
func (entity *TreeEntity) SetLevel(level NodeLevel) {
    entity.Level = level
}
func (entity *TreeEntity) GetParent() TreeEntityIfc {
    return entity.Parent
}
func (entity *TreeEntity) SetParent(parent TreeEntityIfc) {
    entity.Parent = parent.(*TreeEntity)
}
func (entity *TreeEntity) GetParentChain() []TreeEntityIfc {
    return entity.entities2interfaces(entity.ParentChain)
}
func (entity *TreeEntity) SetParentChain(parentChain []TreeEntityIfc) {
    entity.ParentChain = entity.interfaces2entities(parentChain)
}
func (entity *TreeEntity) GetChildren() []TreeEntityIfc {
    return entity.entities2interfaces(entity.Children)
}
func (entity *TreeEntity) SetChildren(children []TreeEntityIfc) {
    entity.Children = entity.interfaces2entities(children)
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

func (entity *TreeEntity) interfaces2entities(ifcs []TreeEntityIfc) []*TreeEntity {
    entities := make([]*TreeEntity, len(ifcs))
    for i, ifc := range ifcs {
        entities[i] = ifc.(*TreeEntity)
    }
    return entities
}
func (entity *TreeEntity) entities2interfaces(entities []*TreeEntity) []TreeEntityIfc {
    ifcs := make([]TreeEntityIfc, len(entities))
    for i, e := range entities {
        ifcs[i] = e
    }
    return ifcs
}
