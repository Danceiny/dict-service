package entity

import (
    . "github.com/Danceiny/dict-service/common"
)

type AreaLevel int

const (
    _ AreaLevel = iota
    CONTINENT
    COUNTRY
    PROVINCE
    CITY
    DISTRICT
    TOWN
    COMMUNITY
)

func (level *AreaLevel) String() string {
    switch *level {
    case 1:
        return "CONTINENT"
    case 2:
        return "COUNTRY"
    case 3:
        return "PROVINCE"
    case 4:
        return "CITY"
    case 5:
        return "DISTRICT"
    case 6:
        return "TOWN"
    case 7:
        return "COMMUNITY"
    }
    return ""
}

type AreaEntity struct {
    TreeEntity
    Pid         NodeId    `gorm:"column:parent_bid" json:"parentBid"`
    Name        string    `gorm:"column:node_name" json:"name"`
    EnglishName string    `gorm:"column:english_name" json:"englishName"`
    Level       AreaLevel `gorm:"column:node_level" json:"level"`
    Code        uint      `gorm:"column:area_code" json:"areaCode"`
}

func (AreaEntity) TableName() string {
    return "dict_area"
}

func (AreaEntity) GetType() DictTypeEnum {
    return AREA
}

func (entity *AreaEntity) GetParentBid() BID {
    return entity.Pid
}

func (entity *AreaEntity) GetDefaultBid() BID {
    // abstract method, implemented by subclass
    return NodeId(-1)
}
