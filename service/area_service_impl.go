package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence"
    . "github.com/Danceiny/dict-service/persistence/entity"
)

var AreaServiceImplCpt *AreaServiceImpl

func init() {
    AreaServiceImplCpt = &AreaServiceImpl{
        TreeSe
        BaseCacheServiceImplCpt}
}

type AreaServiceImpl struct {
    treeService TreeService
    cacheService BaseCacheService
}

func (impl *AreaServiceImpl) GetArea(id NodeId, p, c int) *AreaVO {
    var rawEntity EntityIfc
    var entity *AreaEntity
    rawEntity = impl.cacheService.GetEntityCache(AREA, id, false)
    if rawEntity == nil {
        rawEntity = RepoCpt.Get(AREA, id, false, false)
        if rawEntity == nil {
            return nil
        } else {
            entity = rawEntity.(*AreaEntity)
            impl.cacheService.CacheEntity(AREA, entity, false)
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
