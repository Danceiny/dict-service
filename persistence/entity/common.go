package entity

import (
    "encoding/json"
    . "github.com/Danceiny/dict-service/common"
    "github.com/Danceiny/go.fastjson"
    log "github.com/sirupsen/logrus"
)

func Bids2Json(bids []BID) []byte {
    bytes, err := json.Marshal(bids)
    if err != nil {
        log.Warningf("transfer bids to json error: %v", err)
    }
    return bytes
}

type DynamicAttrPlugin interface {
    GetAttr() *fastjson.JSONObject
    SetAttr(bytes []byte)
}
