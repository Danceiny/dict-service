package service

import (
    . "github.com/Danceiny/dict-service/persistence"
    . "github.com/Danceiny/dict-service/persistence/entity"
)

type AreaService interface {
    GetArea(id int) *AreaVO
}

type AreaServiceImpl struct {
}

func (impl *AreaServiceImpl) GetArea(id int) *AreaVO {
    var rawEntity Entity
    var entity *AreaEntity
    rawEntity = BaseCacheServiceImplCpt.GetEntityCache(AREA, id, false)
    if rawEntity == nil {
        rawEntity = Repo.Get(AREA, id, false, false)
        if rawEntity == nil {
            return nil
        } else {
            entity = rawEntity.(*AreaEntity)
            BaseCacheServiceImplCpt.CacheEntity(AREA, entity, false)
        }
    } else {
        entity = rawEntity.(*AreaEntity)
    }
    return impl.TransferEntity2VO(entity)
}

func (impl *AreaServiceImpl) TransferEntity2VO(entity *AreaEntity) *AreaVO {
    areaVO := &AreaVO{}
    areaVO.Bid = entity.ID
    areaVO.Name = entity.Name
    areaVO.Code = entity.Code
    areaVO.EnglishName = entity.EnglishName
    areaVO.Level = entity.Level
    areaVO.Attr = entity.Attr
    return areaVO
}
