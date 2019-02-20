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
	areaVO.Level = AreaLevel(entity.Level)
	areaVO.Attr = entity.Attr
	areaVO.Pid = entity.Pid
	areaVO.IsCountyCity = entity.IsCountyCity
	var parentChain = entity.ParentChain
	if cap(parentChain) != 0 {
		var parentChainVos = make([]*AreaVO, len(parentChain))
		for i, parent := range parentChain {
			parentChainVos[i] = impl.TransferEntity2VO(parent.(*AreaEntity))
		}
		areaVO.ParentChain = parentChainVos
	}
	var children = entity.Children
	if cap(children) != 0 {
		var childrenVos = make([]*AreaVO, len(children))
		for i, child := range children {
			childrenVos[i] = impl.TransferEntity2VO(child.(*AreaEntity))
		}
		areaVO.Children = childrenVos
	}
	// todo
	areaVO.IsMunicipality = false
	return areaVO
}
