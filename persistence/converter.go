package persistence

import (
    "encoding/json"
    . "github.com/Danceiny/dict-service/persistence/entity"
    log "github.com/sirupsen/logrus"
)

func ParseEntityFromJSON(t DictTypeEnum, bytes []byte) Entity {
    if bytes == nil {
        return nil
    }
    switch t {
    case AREA:
        var entity AreaEntity
        if err := json.Unmarshal(bytes, &entity); err != nil {
            log.Warning(err)
            return nil
        } else {
            return &entity
        }
    }
    return nil
}
