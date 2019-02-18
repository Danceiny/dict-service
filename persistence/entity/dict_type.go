package entity

import (
    "encoding/json"
    . "github.com/Danceiny/dict-service/common"
    "github.com/Danceiny/go.fastjson"
    log "github.com/sirupsen/logrus"
)

type DictTypeEnum int

const (
    _ DictTypeEnum = iota
    CATEGORY
    AREA
    CAR
)

func (t *DictTypeEnum) GetBidColName() string {
    if CATEGORY == *t {
        return "bid"
    } else {
        return "id"
    }
}

func (t *DictTypeEnum) GetTableName() string {
    switch *t {
    case CATEGORY:
        return "dict_category"
    case AREA:
        return "dict_area"
    default:
        return ""
    }
}

func (t *DictTypeEnum) UseHashCache() bool {
    return (*t) == CATEGORY || (*t) == AREA
}

func (t *DictTypeEnum) String() string {
    switch *t {
    case CATEGORY:
        return "CATEGORY";
    case AREA:
        return "AREA";
    case CAR:
        return "CAR"
    }
    return ""
}

func (t *DictTypeEnum) InitEmpty() EntityIfc {
    switch *t {
    case CATEGORY:
        return &CategoryEntity{}
    case AREA:
        return &AreaEntity{}
    case CAR:
        return &CarEntity{}
    }
    return nil
}

func (t *DictTypeEnum) parseBids(bytes []byte, bids *[]BID) {
    return fastjson.ParseArray(string(bytes))
}

func (t *DictTypeEnum)ParseBids(bytes []byte) (bids []BID) {
    t.parseBids(bytes, &bids)
    return bids
}



func (t *DictTypeEnum)MultiParseBids(bytess [][]byte) [][]BID {
    bidss := make([][]BID, len(bytess))
    for i, bytes := range bytess {
        parseBids(bytes, &bidss[i])
    }
    return bidss
}