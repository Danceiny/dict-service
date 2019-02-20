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
	Cids        []BID           `gorm:"-" json:"-"`
	Pids        []BID           `gorm:"-" json:"-"`
	Children    []TreeEntityIfc `gorm:"-" json:"-"`
	Parent      TreeEntityIfc   `gorm:"-" json:"-"`
	ParentChain []TreeEntityIfc `gorm:"-" json:"-"`
	OldLevel    *int            `gorm:"-" json:"-"`
	OldPid      BID             `gorm:"-" json:"-"`
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
