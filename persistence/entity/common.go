package entity

import (
	"encoding/json"
	. "github.com/Danceiny/dict-service/common/FastJson"
	log "github.com/sirupsen/logrus"
)

type BID interface {
	String() string
}

type DynamicAttrPlugin interface {
	GetAttr() JsonObject
	SetAttr(bytes []byte)
}

func ParseBids(bytes []byte) (bids []BID) {
	if err := json.Unmarshal(bytes, &bids); err != nil {
		log.Warningf("parse bids error: %v", err)
	}
	return bids
}
