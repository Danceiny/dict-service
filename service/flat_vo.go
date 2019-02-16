package service

import (
	"github.com/Danceiny/dict-service/common/FastJson"
)

type FlatVO interface {
	ToFlatVO() *FastJson.JsonObject
}
