package persistence

import (
    "encoding/json"
    . "github.com/Danceiny/dict-service/persistence/entity"
    log "github.com/sirupsen/logrus"
)

func ParseEntityFromJSON(t DictTypeEnum, bytes []byte) EntityIfc {
    if bytes == nil {
        return nil
    }
    var err error
    switch t {
    case AREA:
        var entity AreaEntity
        if err = json.Unmarshal(bytes, &entity); err == nil {
            return &entity
        }
        break
    case CATEGORY:
        var entity CategoryEntity
        if err = json.Unmarshal(bytes, &entity); err == nil {
            return &entity
        }
        break
    case CAR:
        var entity CarEntity
        if err = json.Unmarshal(bytes, &entity); err == nil {
            return &entity
        }
        break
    }
    log.Warning(err)
    return nil
}
