package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence/entity"
    "log"
)

var (
    AreaServiceImplCpt *AreaServiceImpl
)

type AreaServiceImpl struct {
    RepoServ  RepositoryService
    TreeServ  TreeService
    CacheServ BaseCacheService
}

func (impl *AreaServiceImpl) GetArea(id NodeId, p, c int) *AreaVO {
    var rawEntity EntityIfc
    var entity *AreaEntity
    rawEntity = impl.CacheServ.GetEntityCache(AREA, id, false)
    if rawEntity == nil {
        rawEntity = impl.RepoServ.Get(AREA, id, false, false)
        if rawEntity == nil {
            log.Println("11111111111231")
            return nil
        } else {
            entity = rawEntity.(*AreaEntity)
            impl.CacheServ.CacheEntity(AREA, entity, false)
        }
    } else {
        entity = rawEntity.(*AreaEntity)
    }
    return impl.TransferEntity2VO(entity)
}

func (impl *AreaServiceImpl) TransferEntity2VO(entity *AreaEntity) *AreaVO {
    if entity == nil {
        return nil
    }
    areaVO := &AreaVO{}
    areaVO.Bid = entity.ID
    areaVO.Name = entity.Name
    areaVO.Code = entity.Code
    areaVO.EnglishName = entity.EnglishName
    areaVO.Level = entity.Level
    areaVO.Attr = entity.Attr
    return areaVO
}
