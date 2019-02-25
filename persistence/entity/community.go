package entity

import (
    "github.com/Danceiny/go.fastjson"
)

type CommunityEntity struct {
    BaseEntity
}

func (CommunityEntity) TableName() string {
    return "dict_community"
}

func (CommunityEntity) GetType() DictTypeEnum {
    return COMMUNITY
}

func (entity *CommunityEntity) ToJSONB() []byte {
    return fastjson.ToJSONB(entity)
}
