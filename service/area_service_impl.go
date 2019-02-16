package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence/entity"
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
    return impl.transferEntity2VO(impl.TreeServ.GetTree(AREA, id, p, c, false, false))
}

func (impl *AreaServiceImpl) TransferEntity2VO(entity *AreaEntity) *AreaVO {
    return impl.transferEntity2VO(entity)
}

func (impl *AreaServiceImpl) transferEntity2VO(ptr interface{}) *AreaVO {
    if ptr == nil {
        return nil
    }
    entity := ptr.(*AreaEntity)
    areaVO := &AreaVO{}
    areaVO.Bid = entity.ID
    areaVO.Name = entity.Name
    areaVO.Code = entity.Code
    areaVO.EnglishName = entity.EnglishName
    areaVO.Level = entity.Level
    areaVO.Attr = entity.Attr
    return areaVO
}
