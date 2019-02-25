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

func (impl *AreaServiceImpl) GetArea(id INT, p, c int) *AreaVO {
    return impl.transferEntity2VO(impl.TreeServ.GetTree(AREA, id, p, c, false, false), false)
}

func (impl *AreaServiceImpl) TransferEntity2VO(entity *AreaEntity, loadBrother bool) *AreaVO {
    return impl.transferEntity2VO(entity, loadBrother)
}

func (impl *AreaServiceImpl) transferEntity2VO(ptr interface{}, loadBrother bool) *AreaVO {
    if ptr == nil {
        return nil
    }
    entity := ptr.(*AreaEntity)
    vo := &AreaVO{}
    vo.Bid = entity.ID
    vo.Name = entity.Name
    vo.Code = entity.Code
    vo.EnglishName = entity.EnglishName
    vo.Level = AreaLevel(entity.Level)
    vo.Attr = entity.Attr
    vo.Pid = entity.Pid
    vo.IsCountyCity = entity.IsCountyCity
    vo.Cids = entity.Cids
    vo.Pids = entity.Pids
    var parentChain = entity.ParentChain
    if cap(parentChain) != 0 {
        var parentChainVos = make([]*AreaVO, len(parentChain))
        for i, parent := range parentChain {
            parentChainVos[i] = impl.transferEntity2VO(
                parent, false)
        }
        vo.ParentChain = parentChainVos
    }
    var children = entity.Children
    if cap(children) != 0 {
        var childrenVos = make([]*AreaVO, len(children))
        for i, child := range children {
            if child != nil {
                childrenVos[i] = impl.transferEntity2VO(
                    child, false)
            }
        }
        vo.Children = childrenVos
    }
    // todo
    vo.IsMunicipality = entity.GetIsMunicipality()
    return vo
}

func (impl *AreaServiceImpl) BatchQuery(bids []BID, simple bool, p int, c int, onlyId bool) []*AreaVO {
    var entities = impl.TreeServ.MultiGet(AREA, bids,
        simple, p, c, onlyId, false, false)
    var ret = make([]*AreaVO, len(bids))
    for i, entity := range entities {
        ret[i] = impl.transferEntity2VO(entity, false)
    }
    return ret
}
