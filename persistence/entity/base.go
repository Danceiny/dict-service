package entity

import (
    . "github.com/Danceiny/dict-service/common"
    "github.com/Danceiny/go.fastjson"
)

type BaseEntity struct {
    ID          NodeId `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
    CreatedTime int64  `gorm:"column:created_time" json:"createdTime"`
    UpdatedTime int64  `gorm:"column:modified_time" json:"modifiedTime"`
    // 不叫DeletedAt，避开gorm的软删除
    DeletedTime int64                `gorm:"column:deleted_time" json:"deletedTime"`
    Attr        *fastjson.JSONObject `gorm:"-" json:"attr"`
}

func (entity *BaseEntity) GetAttr() *fastjson.JSONObject {
    return entity.Attr
}

func (entity *BaseEntity) SetAttr(bytes []byte) {
    o := fastjson.ParseObjectB(bytes)
    entity.Attr = o
}

func (entity *BaseEntity) IsEmpty() bool {
    s := entity.GetBid().String()
    return s == "0" || s == ""
}

func (entity *BaseEntity) GetBid() BID {
    return entity.ID
}
func (entity *BaseEntity) GetCreatedTime() int64 {
    return entity.CreatedTime
}
func (entity *BaseEntity) SetCreatedTime(t int64) {
    entity.CreatedTime = t
}
func (entity *BaseEntity) GetDeletedTime() int64 {
    return entity.DeletedTime
}
func (entity *BaseEntity) SetDeletedTime(t int64) {
    entity.DeletedTime = t
}
func (entity *BaseEntity) GetUpdatedTime() int64 {
    return entity.UpdatedTime
}
func (entity *BaseEntity) SetUpdatedTime(t int64) {
    entity.UpdatedTime = t
}

type EntityIfc interface {
    GetBid() BID
    GetAttr() *fastjson.JSONObject
    SetAttr(bytes []byte)
    GetType() DictTypeEnum
    IsEmpty() bool
    GetCreatedTime() int64
    SetCreatedTime(int64)
    GetDeletedTime() int64
    SetDeletedTime(int64)
    GetUpdatedTime() int64
    SetUpdatedTime(int64)
}
