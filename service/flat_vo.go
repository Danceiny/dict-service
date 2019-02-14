package service

import (
    "github.com/Danceiny/dict-service/common/FastJson"
)

type FlatVO interface {
    toFlatVO() FastJson.JsonObject
}
