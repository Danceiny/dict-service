package persistence

import (
    . "github.com/Danceiny/dict-service/persistence/entity"
)

func ParseEntityFromJSON(t DictTypeEnum, bytes []byte) EntityIfc {
    return t.ParseJSONB(bytes)
}
