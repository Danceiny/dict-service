package service

import (
	FastJson "github.com/Danceiny/go.fastjson"
)

type FlatVO interface {
	toFlatVO() FastJson.JSONObject
}
