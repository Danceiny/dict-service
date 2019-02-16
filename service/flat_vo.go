package service

import (
	"github.com/Danceiny/go.fastjson"
)

type FlatVO interface {
	ToFlatVO() *fastjson.JSONObject
}
