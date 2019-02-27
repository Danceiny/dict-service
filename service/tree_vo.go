package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/go.fastjson"
)

type TreeVO struct {
    Name      string      `json:"name"`
    Weight    int         `json:"weight"`
    LevelName string      `json:"levelName"`
    Cids      []BID       `json:"childrenBidList"`
    Pids      []BID       `json:"parentChainBidList"`
    Attr      *JSONObject `json:"attr"`
}
