package entity

import (
    "encoding/json"
    . "github.com/Danceiny/dict-service/common"
    "github.com/Danceiny/go.fastjson"
    "github.com/sirupsen/logrus"
)

type DictTypeEnum int

const (
    COMMON DictTypeEnum = iota
    CATEGORY
    AREA
    CAR
)

func (t DictTypeEnum) GetBidColName() string {
    if CATEGORY == t {
        return "bid"
    } else {
        return "id"
    }
}

func (t DictTypeEnum) GetTableName() string {
    switch t {
    case CATEGORY:
        return "dict_category"
    case AREA:
        return "dict_area"
    case CAR:
        return "dict_car"
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
        return "CATEGORY"
    case AREA:
        return "AREA"
    case CAR:
        return "CAR"
    }
    return ""
}

func (t DictTypeEnum) InitEmpty() EntityIfc {
    switch t {
    case CATEGORY:
        return &CategoryEntity{}
    case AREA:
        return &AreaEntity{}
    case CAR:
        return &CarEntity{}
    }
    return nil
}
func (t DictTypeEnum) ParseBidsB(bytes []byte) []BID {
    if bytes == nil {
        return nil
    }
    return t.ParseBids(string(bytes))
}

func (t DictTypeEnum) ParseBids(bytes string) []BID {
    ja := fastjson.ParseArray(bytes)
    if ja == nil {
        return nil
    }
    bids := make([]BID, 0, ja.Size())
    switch t {
    case CATEGORY:
        for ja.Next() {
            bids = append(bids, STRING(ja.Current().(string)))
        }
        break
    default:
        for ja.Next() {
            bids = append(bids, INT(ja.Current().(float64)))
        }
        break
    }
    return bids
}

func (t DictTypeEnum) MultiParseBids(jsons []interface{}) [][]BID {
    var bidss = make([][]BID, len(jsons))
    for i, bytes := range jsons {
        if bytes != nil {
            bidss[i] = t.ParseBids(bytes.(string))
        }
    }
    return bidss
}
func (t DictTypeEnum) ParseJSON(s string) EntityIfc {
    if s == "" {
        return nil
    }
    var entity = t.InitEmpty()
    fastjson.ParseObjectT(s, entity)
    return entity
}

func (t DictTypeEnum) ParseJSONB(bytes []byte) EntityIfc {
    if bytes == nil {
        return nil
    }
    var err error
    switch t {
    case AREA:
        var entity AreaEntity
        if err = json.Unmarshal(bytes, &entity); err == nil {
            return &entity
        }
        break
    case CATEGORY:
        var entity CategoryEntity
        if err = json.Unmarshal(bytes, &entity); err == nil {
            return &entity
        }
        break
    case CAR:
        var entity CarEntity
        if err = json.Unmarshal(bytes, &entity); err == nil {
            return &entity
        }
        break
    }
    logrus.Warningf("ParseJSONB err: %v", err)
    return nil
}
