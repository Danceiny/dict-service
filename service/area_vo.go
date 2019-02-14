package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/common/FastJson"
    . "github.com/Danceiny/dict-service/persistence/entity"
)

type AreaVO struct {
    Bid         NodeId      `json:"bid"`
    Name        string      `json:"name"`
    EnglishName string      `json:"englishName"`
    Level       AreaLevel   `json:"-"`
    LevelName   string      `json:"levelName"`
    Code        uint        `json:"areaCode"`
    Attr        *JsonObject `json:"attr"`
}

func (vo *AreaVO) ToFlatVO() *JsonObject {
    attr := vo.Attr
    if attr == nil || *attr == nil {
        attr = &JsonObject{}
    }
    attr.PutFluent("bid", vo.Bid).
        PutFluent("name", vo.Name).
        PutFluent("englishName", vo.EnglishName).
        PutFluent("levelName", vo.Level.String()).
        PutFluent("areaCode", vo.Code)
    return attr
}

type PagableAreaVO struct {
    Total int      `json:"total"`
    Page  int      `json:"page"`
    List  []AreaVO `json:"list"`
}
