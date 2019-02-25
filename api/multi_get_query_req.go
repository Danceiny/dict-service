package api

import (
    . "github.com/Danceiny/dict-service/common"
)

type MultiGetQueryReq struct {
    Category    []STRING `json:"category"`
    Unknown     []STRING `json:"unknown"`
    Area        []INT    `json:"area"`
    Car         []INT    `json:"car"`
    Community   []INT    `json:"community"`
    ParentDepth int      `json:"parentDepth"`
    HasChildren int      `json:"hasChildren"`
    OnlyId      bool     `json:"onlyId"` // 默认false
}
