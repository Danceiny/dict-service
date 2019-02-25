package entity

import (
    . "github.com/Danceiny/dict-service/common"
    "github.com/Danceiny/go.fastjson"
)

type AreaLevel int

const (
    CONTINENT AreaLevel = iota
    COUNTRY
    PROVINCE
    CITY
    DISTRICT
    TOWN
    // COMMUNITY
)

var MunicipalityIds = []int{28, 30, 31, 32}

func (level AreaLevel) String() string {
    switch level {
    case CONTINENT:
        return "CONTINENT"
    case COUNTRY:
        return "COUNTRY"
    case PROVINCE:
        return "PROVINCE"
    case CITY:
        return "CITY"
    case DISTRICT:
        return "DISTRICT"
    case TOWN:
        return "TOWN"
        // case COMMUNITY:
        //     return "COMMUNITY"
    }
    return ""
}

func (level AreaLevel) Val() int {
    return int(level)
}

type AreaEntity struct {
    TreeEntity
    Pid            INT    `gorm:"column:parent_bid" json:"parentBid"`
    Name           string `gorm:"column:node_name" json:"name"`
    EnglishName    string `gorm:"column:english_name" json:"englishName"`
    Code           INT    `gorm:"column:area_code" json:"areaCode"`
    IsCountyCity   bool   `gorm:"isCountyCity" json:"isCountyCity"`
    IsMunicipality bool   `gorm:"-" json:"IsMunicipality"`
}

func (AreaEntity) TableName() string {
    return "dict_area"
}

func (AreaEntity) GetType() DictTypeEnum {
    return AREA
}

func (entity *AreaEntity) GetPid() BID {
    return entity.Pid
}

func (entity *AreaEntity) GetDefaultBid() BID {
    // abstract method, implemented by subclass
    return INT(-1)
}

func (entity *AreaEntity) GetIsMunicipality() bool {
    var id = int(entity.ID)
    for _, v := range MunicipalityIds {
        if id == v {
            return true
        }
    }
    return false
}

func (entity *AreaEntity) ToJSONB() []byte {
    entity.IsMunicipality = entity.GetIsMunicipality()
    return fastjson.ToJSONB(entity)
}
