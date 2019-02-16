package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence/entity"
    . "github.com/Danceiny/go.fastjson"
    "log"
)

type AreaVO struct {
    Bid         NodeId      `json:"bid"`
    Pid         NodeId      `json:"pid"`
    Name        string      `json:"name"`
    EnglishName string      `json:"englishName"`
    Level       AreaLevel   `json:"-"`
    LevelName   string      `json:"levelName"`
    Code        uint        `json:"areaCode"`
    Attr        *JSONObject `json:"attr"`
}

func (vo *AreaVO) ToFlatVO() *JSONObject {
    if nil == vo {
        log.Println("`````````")
        return nil
    }
    attr := vo.Attr
    if attr == nil || *attr == nil {
        attr = &JSONObject{}
    }
    attr.PutFluent("bid", vo.Bid).
        PutFluent("pid", vo.Pid).
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
