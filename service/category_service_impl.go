package service

import (
    . "github.com/Danceiny/dict-service/common"
)

var (
    CategoryServiceImplCpt *CategoryServiceImpl
)

type CategoryServiceImpl struct {
}

func (impl *CategoryServiceImpl) BatchQuery(ids []BID, simple bool, p int, c int, onlyId bool) []*CategoryVO {
    return nil
}
