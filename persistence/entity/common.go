package entity

import (
	"encoding/json"
	"github.com/Danceiny/go.fastjson"
	log "github.com/sirupsen/logrus"
)

type BID interface {
	String() string
}

func ParseBids(bytes []byte) (bids []BID) {
	if err := json.Unmarshal(bytes, &bids); err != nil {
		log.Warningf("parse bids error: %v", err)
	}
	return bids
}

type DynamicAttrPlugin interface {
	GetAttr() *fastjson.JSONObject
	SetAttr(bytes []byte)
}
