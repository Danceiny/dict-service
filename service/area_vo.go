package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence/entity"
    . "github.com/Danceiny/go.fastjson"
)

type AreaVO struct {
    Bid            INT         `json:"bid"`
    Pid            INT         `json:"pid"`
    Name           string      `json:"name"`
    EnglishName    string      `json:"englishName"`
    Level          AreaLevel   `json:"-"`
    LevelName      string      `json:"levelName"`
    Code           INT         `json:"areaCode"`
    Children       []*AreaVO   `json:"children"`
    Cids           []BID       `json:"childrenBidList"`
    ParentChain    []*AreaVO   `json:"parentChain"`
    Pids           []BID       `json:"parentChainBidList"`
    Brothers       []*AreaVO   `json:"brothers"`
    IsMunicipality bool        `json:"isMunicipality"`
    IsCountyCity   bool        `json:"isCountyCity"`
    Attr           *JSONObject `json:"attr"`
}

func (vo *AreaVO) ToFlatVO() *JSONObject {
    if nil == vo {
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
        PutFluent("areaCode", vo.Code).
        PutFluent("children", vo.Children).
        PutFluent("childrenBidList", vo.Cids).
        PutFluent("parentChain", vo.ParentChain).
        PutFluent("parentChainBidList", vo.ParentChain)
    return attr
}

type PagableAreaVO struct {
    Total int      `json:"total"`
    Page  int      `json:"page"`
    List  []AreaVO `json:"list"`
}
