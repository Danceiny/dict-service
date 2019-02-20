package entity

import (
	. "github.com/Danceiny/dict-service/common"
	"github.com/Danceiny/go.fastjson"
)

type DictTypeEnum int

const (
	COMMON DictTypeEnum = iota
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

func (t *DictTypeEnum) ParseBids(bytes []byte) []BID {
	ja := fastjson.ParseArrayB(bytes)
	if ja == nil {
		return nil
	}
	bids := make([]BID, ja.Size())
	switch *t {
	case CATEGORY:
		for ja.Next() {
			bids = append(bids, STRING(ja.Current().(string)))
		}
		break
	default:
		for ja.Next() {
			bids = append(bids, INT(ja.Current().(int64)))
		}
		break
	}
	return bids
}

func (t *DictTypeEnum) MultiParseBids(bytess [][]byte) [][]BID {
	bidss := make([][]BID, len(bytess))
	for i, bytes := range bytess {
		bidss[i] = t.ParseBids(bytes)
	}
	return bidss
}
